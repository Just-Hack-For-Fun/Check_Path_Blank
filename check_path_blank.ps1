param(
    [Parameter(ValueFromRemainingArguments = $true)]
    [string[]]$Paths
)

$startTime = Get-Date

# 如果没有参数，获取所有本地盘符
if (-not $Paths -or $Paths.Count -eq 0) {
    $Paths = (Get-PSDrive -PSProvider 'FileSystem' | Where-Object { $_.Root -match '^[A-Z]:\\$' }).Root
}

# 获取可执行扩展名数组
$pathexts = $env:PATHEXT.ToLower().Split(';') | ForEach-Object { $_.TrimStart('.') } | Where-Object { $_ }

Write-Host "`n检测带空格路径的截断前缀可执行文件风险" -ForegroundColor Cyan

foreach ($rootPath in $Paths) {
    Write-Host "`n扫描路径: $rootPath" -ForegroundColor Gray

    # 递归扫描所有文件和文件夹
    Get-ChildItem -LiteralPath $rootPath -Recurse -Force -ErrorAction SilentlyContinue |
        Where-Object { $_.FullName -match '\s' } | # 路径含空格
        ForEach-Object {
            $fullPath = $_.FullName
            $baseRoot = $rootPath
            if ($fullPath.StartsWith($baseRoot)) {
                $relativePath = $fullPath.Substring($baseRoot.Length)
            } else {
                $relativePath = $fullPath
            }

            if (-not $_.PSIsContainer) {
                $spaceIndexes = (0..($relativePath.Length - 1)) | Where-Object { $relativePath[$_] -eq ' ' }
                foreach ($spaceIdx in $spaceIndexes) {
                    $prefix = $relativePath.Substring(0, $spaceIdx)
                    $checkBase = Join-Path $baseRoot $prefix

                    if (Test-Path $checkBase -PathType Leaf) {
                        Write-Host "发现可疑文件: $checkBase (截断自 $fullPath)" -ForegroundColor Yellow
                    }

                    foreach ($ext in $pathexts) {
                        $exePath = "$checkBase.$ext"
                        if (Test-Path $exePath -PathType Leaf) {
                            Write-Host "发现可疑可执行文件: $exePath (截断自 $fullPath)" -ForegroundColor Red
                        }
                    }
                }
            }
        }
}

$endTime = Get-Date
$duration = $endTime - $startTime

Write-Host "`n扫描开始时间: $startTime"
Write-Host "扫描结束时间: $endTime"
Write-Host ("总耗时: {0:hh\:mm\:ss}" -f $duration)
