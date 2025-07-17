# Check_Path_Blank
检查 Windows 平台上存在空格的路径是否存在可能的劫持文件的脚本

在 Windows 系统中，部分程序在解析程序路径时，遇到存在空格的路径，如果没有使用双引号包裹，可能存在截断的问题，这个问题我们在 Windows 服务的文章中已经有提到，但是我发现不同程序的处理是不一样的，因此我们采取相对严格的排查策略

如果系统中存在 `C:\t m p\a b c.exe` ，那么我们应该检查以下文件是否存在

- `C:\t`
- `C:\t.com`
- `C:\t.exe`

- 其他后缀，具体根据 pathext 环境变量
- `C:\t m`
- `C:\t m.com`
- `C:\t m.exe`
- 其他后缀，具体根据 pathext 环境变量
- `C:\t m p\a`
- `C:\t m p\a.com`
- `C:\t m p\a.exe`
- 其他后缀，具体根据 pathext 环境变量
- `C:\t m p\a b`
- `C:\t m p\a b.com`
- `C:\t m p\a b.exe`
- 其他后缀，具体根据 pathext 环境变量



下面提供 PowerShell 和 Go 语言两个版本的排查脚本



#### PowerShell

```bat
check_path_blank.ps1 <path>
```

可以指定检查路径，例如 `C:\` ，如果不加参数。默认会检查所有的盘符的所有路径，包括共享盘符


#### Golang

```bat
# 编译
go build -o check_path_blank.exe check_path_blank.go

# 使用
check_path_blank.exe <path>
```

可以指定检查路径，例如 `C:\` ，如果不加参数。默认会检查所有的盘符的所有路径，包括共享盘符

Golang 语言的程序可能比 PowerShell 效率更高一些







------

【 Windows Server 2016 】默认情况

![image-20250717215815022](http://mweb-tc.oss-cn-beijing.aliyuncs.com/2025-07-17-135817.png)

```
检测带空格路径的截断前缀可执行文件风险

扫描路径: C:\
发现可疑文件: C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps (截断自 C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps Icons)
发现可疑文件: C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps (截断自 C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps Icons-journal)
发现可疑文件: C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 (截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64)
发现可疑文件: C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 (截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64 Critical)
发现可疑文件: C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64 (截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64 Critical)
发现可疑文件: C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 (截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 Critical)
发现可疑文件: C:\Windows\System32\Tasks\Microsoft\Windows\Data Integrity Scan\Data Integrity Scan (截断自 C:\Windows\System32\Tasks\Microsoft\Windows\Data Integrity Scan\Data Integrity Scan for Crash Recovery)
发现可疑可执行文件: C:\Windows\WinSxS\amd64_microsoft-windows-iis-legacysnapin_31bf3856ad364e35_10.0.14393.0_none_ae953f82c8b8a231\IIS6.msc (截断自 C:\Windows\WinSxS\amd64_microsoft-windows-iis-legacysnapin_31bf3856ad364e35_10.0.14393.0_none_ae953f82c8b8a231\IIS6 Manager.lnk)
发现可疑可执行文件: C:\Windows\WinSxS\amd64_microsoft-windows-iis-managementconsole_31bf3856ad364e35_10.0.14393.0_none_b54808dbd1ca2029\IIS.msc (截断自 C:\Windows\WinSxS\amd64_microsoft-windows-iis-managementconsole_31bf3856ad364e35_10.0.14393.0_none_b54808dbd1ca2029\IIS Manager.lnk)
```





![image-20250717215720163](http://mweb-tc.oss-cn-beijing.aliyuncs.com/2025-07-17-135721.png)

```
检测带空格路径的截断前缀可执行文件风险
开始扫描，请耐心等待... 开始时间: 2025-07-17 21:55:58

扫描路径: C:\
[可疑文件] C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps （截断自 C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps Icons）
[可疑文件] C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps （截断自 C:\Users\Administrator\AppData\Local\Microsoft\Edge\User Data\Default\HubApps Icons-journal）
[可疑文件] C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 （截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64）
[可疑文件] C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 （截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64 Critical）
[可疑文件] C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64 （截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 64 Critical）
[可疑文件] C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 （截断自 C:\Windows\System32\Tasks\Microsoft\Windows\.NET Framework\.NET Framework NGEN v4.0.30319 Critical）
[可疑文件] C:\Windows\System32\Tasks\Microsoft\Windows\Data Integrity Scan\Data Integrity Scan （截断自 C:\Windows\System32\Tasks\Microsoft\Windows\Data Integrity Scan\Data Integrity Scan for Crash Recovery）
[可疑可执行文件] C:\Windows\WinSxS\amd64_microsoft-windows-iis-legacysnapin_31bf3856ad364e35_10.0.14393.0_none_ae953f82c8b8a231\IIS6.msc （截断自 C:\Windows\WinSxS\amd64_microsoft-windows-iis-legacysnapin_31bf3856ad364e35_10.0.14393.0_none_ae953f82c8b8a231\IIS6 Manager.lnk）
[可疑可执行文件] C:\Windows\WinSxS\amd64_microsoft-windows-iis-managementconsole_31bf3856ad364e35_10.0.14393.0_none_b54808dbd1ca2029\IIS.msc （截断自 C:\Windows\WinSxS\amd64_microsoft-windows-iis-managementconsole_31bf3856ad364e35_10.0.14393.0_none_b54808dbd1ca2029\IIS Manager.lnk）
```





> 参考文章：
> https://mp.weixin.qq.com/s/_OLwgWbrnAhXLGdc0n_Kaw

