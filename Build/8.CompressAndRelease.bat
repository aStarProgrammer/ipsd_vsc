cd ..\..\ipsd_vsc_release
set ipsdVersion=0.2.1.1
mkdir %ipsdVersion%

cd Windows
bandizip c ipsd_vsc_%ipsdVersion%_Windows_X64.zip Windows_X64
bandizip c ipsd_vsc_%ipsdVersion%_Windows_X86.zip Windows_X86 

::ping 127.0.0.1 -n 15 > nul 

move /Y ipsd_vsc_%ipsdVersion%_Windows_X64.zip ..\%ipsdVersion%
move /Y ipsd_vsc_%ipsdVersion%_Windows_X86.zip ..\%ipsdVersion%

cd ..
cd Linux

bandizip c ipsd_vsc_%ipsdVersion%_Linux_X64.tgz Linux_X64
bandizip c ipsd_vsc_%ipsdVersion%_Linux_X86.tgz Linux_X86 

::ping 127.0.0.1 -n 15 > nul 

move /Y ipsd_vsc_%ipsdVersion%_Linux_X64.tgz ..\%ipsdVersion%
move /Y ipsd_vsc_%ipsdVersion%_Linux_X86.tgz ..\%ipsdVersion%

cd ..
bandizip c ipsd_vsc_%ipsdVersion%_Darwin64.tgz Darwin64

::ping 127.0.0.1 -n 15 > nul 

move /Y ipsd_vsc_%ipsdVersion%_Darwin64.tgz %ipsdVersion%

bandizip c ipsd_vsc_%ipsdVersion%_Arm6.tgz Arm6

::ping 127.0.0.1 -n 15 > nul 

move /Y ipsd_vsc_%ipsdVersion%_Arm6.tgz %ipsdVersion%

cd ..\ipsd_vsc\Build