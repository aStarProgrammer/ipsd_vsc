# IPSD_VSC (vsc version IPSD)
IPSD_VSC(Inter Planet Site Watchdog for visual studio code extension) is a tool to work with IPSC_VSC and create static html site automatically.

## Background

IPFS (Inter Planet File System [IPFS](https://ipfs.io)) is a peer-to-peer hyperlink protocol which is used to publish content. We can publish a web site  on IPFS as we publish a site on http.

But as IPFS is an p2p system, file published on IPFS cannot be changed, if we changed a file and publish to IPFS again, it is a completely new file from the old one.  Changing files of a IPFS file is not encouraged. So generally sites that are built on ASP.NET Java PHP which have a lot of scripts are not the best option when you want to publish a site to IPFS. Static website based on HTML and CSS is the best option.

IPSC_VSC is the tool to create static html site that you can publish to IPFS.

IPSD_VSC work with IPSC_VSC to create site automatically.

**NOTE:**

This version of ipsd_vsc is used for visual studio code extension, its name is ipsd_vsc, it is almost the same with ipsd_vsc except the method to add meta data of markdown and html source file. 

### If you want to use this tool as IPSD_VSC, do as follows

## Install

Download the release for your platform from Releases, unzip it.


## Build
If you can not find a release for your platform, build it from source code as follows:

1. Install go

2. Install git
   
       	Download and install
       		https://git-scm.com/download
       	OR
       		sudo apt-get install git	

3. Install mingw(Windows)

4. Install Liteide (https://github.com/visualfc/liteide)


   ​	*Windows/Linux/MacOSX just download and install

   ​	*Raspbian

   ​		Download source (qt4 Linux 64)code and compile as follows:

   ​		

   ```bash
       sudo apt-get update
       sudo apt-get upgrade
       sudo apt-get install git
       git clone https://github.com/visualfc/liteide.git
       sudo apt-get install qt4-dev-tools libqt4-dev libqtcore4 libqtgui4 libqtwebkit-dev g++
       cd liteide/build
       ./update_pkg.sh
       export QTDIR=/usr
       ./build_linux.sh
       cd ~/liteide/liteidex
       ./linux_deploy.sh
       cd ~/liteide/liteidex/liteide/bin 
       ./liteide
   ```

5. Open ipsd_vsc with liteide 

6. Select the platform you needed, modify current environment according to step 1 and 3
    Modify GOROOT and PATH

7. Compile->Build

##### NOTE: Run this tool with Administrator Permission

## Commands

* New Monitor 

```bash
ipsd_vsc -Command NewMonitor -SiteFolder -SiteTitle  -MonitorFolder
```

Create a new monitor, connect a originail source folder to a ipsc_vsc site project folder

Example:

```bash
ipsd_vsc -Command NewMonitor -SiteFolder "F:\TestSite" -SiteTitle "Test Site" -MonitorFolder "F:\WatchdogSpace"
```

* Run Monitor

```bash
ipsd_vsc -Command RunMonitor -MonitorFolder -IndexPageSize
```

Run the monitor defined in MonitorFolder , if there are any change in the monitor folder (add delete or update), will update the changes to ispc and then compile site with IndexPageSize

IndexPageSize (for index page and more page of site, for more information, read QuickHelp.txt of FullHelp.txt of ipsc_vsc)

*  Normal 	index(more) page will contain 20 items

*  Small  	index(more) page will contain 10 items

*  VerySmall	index(more) page will contain 5  items

*  Big		index(more) page will contain 30 items

Example:

```bash
ipsd_vsc -Command RunMonitor -MonitorFolder "F:\WatchdogSpace" -IndexPageSize "VerySmall"
```

* List Normal File

```bash
ipsd_vsc -Command ListNormalFile -MonitorFolder 
```

List all the normal files that already added to the connected site project
	
Example:

```bash
ipsd_vsc -Command ListNormalFile -MonitorFolder "F:\WatchdogSpace"
```

## Build Working Environment

IPSD_VSC must work with IPSC_VSC , you need to do as following to create a environment to create static site.

1. Download IPSD_VSC

2. Download IPSC_VSC

3. Unzip IPSD_VSC

4. Unzip IPSC_VSC

5. Copy all the files to IPSD_VSC folder

6. Add this path to PATH environment variable

7. Install pandoc

   https://pandoc.org

   Install it as following KB :

    https://pandoc.org/installing.html 

   Open command (cmd for Windows, shell for Linux/macOS) , run pandoc -v ,it should return the version information

8. Install a markdown editor , Typora or Visual Studio Code (need to install markdown extension) recommended. 

   https://typeora.io

   https://code.visualstudio.com 

9. Create ipsc_vsc source folder(For Example F:\TestSite) and ipsc_vsc output folder(F:\SiteOutputFolder), then use IPSC_VSC to create a ipsc_vsc site project

  ```bash
    ipsc_vsc -Command "NewSite" -SiteFolder "F:\TestSite" -SiteTitle "Test Site" -SiteAuthor "Chao(sdxianchao@gmail.com)" -SiteDescription "Test Site for IPSC_VSC" -OutputFolder "F:\SiteOutputFolder"
  ```

 Note: run this command with administrator permission (windows:Run as Administrator, Linux/Mac sudo), because ipsc_vsc need to call makesoftlink function when it create output folder, and this function need administrator permission.

 You can also don't use OutputFolder argument, then there will be an output folder under site folder created(F:\TestSite\output)

  ```bash
    ipsc_vsc -Command "NewSite" -SiteFolder "F:\TestSite" -SiteTitle "Test Site" -SiteAuthor "Chao(sdxianchao@gmail.com)" -SiteDescription "Test Site for IPSC_VSC"
  ```

  Note: if the command not works as you expected, check the ipsc_vsc.log in the folder that ipsc_vsc locates.

10. Create a empty original source folder for IPSD_VSC. Now there are three folders

  * Orignial Folder F:\WatchdogSpace which contains the orginal files for the site
  * Site Source Folder F:\TestSite which contains files for the ipsc_vsc site
  * Site Output Folder F:\SiteOutputFolder which contains generated files of site

11. Connect Original source folder with IPSC_VSC by ipsd_vsc

```bash
ipsd_vsc -Command NewMonitor -SiteFolder "F:\TestSite" -SiteTitle "Test Site" -MonitorFolder "F:\WatchdogSpace"
```

Note: if the command not works as you expected, check the ipsd_vsc.log in the folder that ipsc_vsc locates.

## How it works

There are a original source folder which contains source files, it looks like

* Files    // The files contained in this folder will be added to IPSC_VSC as normal file

  * SubFolder1
    * SubFile1
    * SubFile2
  * SubFolder2
  * File1
  * File2

* Html  //The files contained in this folder will be added to ipsc_vsc as HTML file

  * H1.html
  * H1.mta.json //H1.mta.json is the file including metadata for H1.html
  * H1.png  //xx.png will be used as title iamge for xx.html
  * H2.html
  * H2.mta.json
  * H2.png

* Link 

  * L1.lik.json //lik.json include the metadata(definition) of a link
  * L2.lik.json
  * L1.png //xx.png will be used as title image for link xx
  * L2.png

* Markdown //Markdown files in this folder will be added to IPSC_VSC as markdown file

  * M1.md 
  * M1.mta.json
  * M1.png //xx.png will be used as title image for xx.md
  * M2.md
  * M2.mta.json
  * M2.png

* monitor.sm  //definition file

  

You can add/remove/update  md,html and link in this folder,  add/remove/update their metadata file or title iamge. Before that, you need to connect this folder with a ipsc_vsc site project. After you modified the orginial folder,call ipsd_vsc command RUNMONITOR, ipsd_vsc will check this folder and send all the changes(Add/Remove/Update) to IPSC_VSC, and call ipsc_vsc to compile the site again. 

## Edit and generate site

#### 1. Add File 

* #### Markdown File

1. Open original source folder->Templates Folder. 

2. Create and edit a markdown file, then save it.

3. Create a mta.json file with the same name, add the following content

     {"Title":"Download ipsc","Author":"chao","Description":"","IsTop":false}

4. Save the file

5. Add a png image with the same name, and its size should be smaller than 30 KB

6. Copy md, mta.json and png files to the Original Source Folder->Markdown



* #### Html File

1. Create or copy a new Html File

2. Create a mta.json with same name and add following content 

     {"Title":"Test Html","Author":"Chao","Description":"","IsTop":true}

3. Save the file

4. Create or copy a png image with the same name, and its size should be smaller than 30 KB

5. Copy html mta.json and png files to the Original Source Folder->Html

* #### Link

 1.Create a lik.json and add content as follows
	{"Url":"https://www.google.com","Title":"google","Author":"chao-PC\\chao","Description":"","IsTop":true}

 2.Save it with name google.lik.json

 3. Add a image named google, such as google.png google.jpg 

 4. Copy the lik.json and image to the Original Source Folder->Link


* #### Normal File

  Copy a file or folder to Originial Source Folder->Files

### 2. Run Monitor

Run ipsd_vsc RunMonitor command, this command will push changes from orignial folder to ipsc_vsc source folder,then call ipsc_vsc to compile the site again, generate new site in ipsc_vsc output folder.

```bash
ipsd_vsc -Command RunMonitor -MonitorFolder "F:\WatchdogSpace" -IndexPageSize "VerySmall"
```

If you already monitored the ipsc_vsc folder with ipsp, ipsp will publish the new site to ipfs.

If you decided to publish the generated site to web server, now you can check the site with browser.

## Publish site to IPFS

1. Download IPSP and unzip it

2. Use IPSP to monitor IPSC_VSC output folder

```bash
ipsp -SiteFolder "F:\TestSite" -MonitorInterval XXX
```

3. When you generate site and the file in output folder updated, ipsp will publish it to ipfs after MonitorInterval



## Raise A Issue

Send email to sdxianchao@gmail.com 



## Maintainers

[@aStarProgrammer](https://github.com/aStarProgrammer).


## License

[MIT](LICENSE)

## HomePage

* Github

  https://github.com/astarprogrammer/ipsc

* IPFS

  http://localhost:8080/ipns/QmYY127PK6pczLrEB1p1mijTFr8RsvRqKFX5q4XepxS1fd/

​	Visit the following page for how to connect ipfs network and visit the above web site

​		https://ipfs.io