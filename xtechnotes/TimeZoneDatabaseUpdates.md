# Updating Go Installations with the Latest IANA Time Zone Data



## Table Of Contents

  + [Introduction](#introduction)
    
    - [IANA Time Zone Information](#iana-time-zone-information)
    
    - [**Go** Relies on the IANA Time Zone Database](#go-relies-on-the-iana-time-zone-database)
    
    - [How Time Zones Are Utilized in **Go** Source Code](#how-time-zones-are-utilized-in-go-source-code)
      
      + [Load Location](#load-location)
      + [Time Package References](#time-package-references)
      
    - [_**Go**_ Provides for IANA Database Updates](#_go_-provides-for-iana-database-updates)
    
      
    
  + [Existing Go Time Zone Configuration](#existing-go-time-zone-configuration)
    - [Configuration Overview](#configuration-overview)
      + [Directory/Folder Structure](#directoryfolder-structure)
      + [The Go Time Directory](#the-go-time-directory)
      + [zoneinfo.zip](#zoneinfozip)
      + [update.bash](#updatebash)
      
    - [Upgrading Go Time Zone Configuration](#upgrading-go-time-zone-configuration)
    
      
    
  + [Objective - Create New File: **zoineinfo.zip**](#objective---create-new-file-zoineinfozip)

    

  + [Two Methods for Updating Go Time Zone Data](#two-methods-for-updating-go-time-zone-data)
    - [Method-1 The Easy Way: **update.bash**](#method-1-the-easy-way-updatebash)
      
      + [Method-1 Step-1 Safety First: Make a Backup Copy of **zoneinfo.zip**](#method-1-step-1-safety-first-make-a-backup-copy-of-zoneinfozip)
      + [Method-1 Step-2 Identify the Desired IANA Time Zone Data Version](#method-1-step-2-identify-the-desired-iana-time-zone-data-version)
        - [Method-1 Access the IANA Time Zone Website](#method-1-access-the-iana-time-zone-website)
        - [Method-1 Record the Target Time Zone Version Number](#method-1-record-the-target-time-zone-version-number)
      + [Method-1 Step-3 Create a Scratch Directory](#method-1-step-3-create-a-scratch-directory)
      + [Method-1 Step-4 Copy *update.bash* to Scratch Directory](#method-1-step-4-copy-updatebash-to-scratch-directory)
      + [Method-1 Step-5 Modify the Bash File](#method-1-step-5-modify-the-bash-file)
      + [Method-1 Step-6 Navigate to Scratch Directory](#method-1-step-6-navigate-to-scratch-directory)
      + [Method-1 Step-7 Execute the Bash Script](#method-1-step-7-execute-the-bash-script)
      + [Method-1 Step-8 Close the Terminal Windows](#method-1-step-8-close-the-terminal-windows)
      + [Method-1 Step-9 Examine Scratch Directory Contents](#method-1-step-9-examine-scratch-directory-contents)
      + [Method-1 Step-10 Copy **zoneinfo.zip** to the **Go** Time Directory](#method-1-step-10-copy-zoneinfozip-to-the-go-time-directory)
      + [Method-1 Step-11 All Done - Mission Accomplished](#method-1-step-11-all-done---mission-accomplished)     
      + [Method-1: Pro’s and Con’s](#method-1-pros-and-cons)
        - [Method-1 Advantages](#method-1-advantages)
        
        - [Method-1 Drawbacks](#method-1-drawbacks)
    - [Method-2: The Not So Easy Way](#method-2-the-not-so-easy-way)
      
      + [Method-2 Step-1 Safety First: Make a Backup Copy of **zoneinfo.zip**](#method-2-step-1-safety-first-make-a-backup-copy-of-zoneinfozip)
      
      + [Method-2 Step-2 Identify the Desired IANA Time Zone Data Version](#method-2-step-2-identify-the-desired-iana-time-zone-data-version)
        
          - [Method-2 Access the IANA Time Zone Website](#method-2-access-the-iana-time-zone-website)
          
          - [Method-2 Record the Target Time Zone Version Number](#method-2-record-the-target-time-zone-version-number)
        
      + [Method-2 Step-3 Create a Scratch Directory](#method-2-step-3-create-a-scratch-directory)
      
      + [Method-2 Step-4 Navigate to Scratch Directory](#method-2-step-4-navigate-to-scratch-directory)
      
      + [Method-2 Step-5 Execute Bash Commands](#method-2-step-5-execute-bash-commands)
        - [Execute Bash Command \#1 - Set Environment to Exit on Error](#execute-bash-command-\1---set-environment-to-exit-on-error)
        - [Execute Bash Command \#2 - Delete Old **work** directory tree](#execute-bash-command-\2---delete-old-work-directory-tree)
        - [Execute Bash Command \#3 - Create **work** Directory](#execute-bash-command-\3---create-work-directory)
        - [Execute Bash Command \#4 - Change to **work** Directory](#execute-bash-command-\4---change-to-work-directory)
        - [Execute Bash Command \#5 - Create **zoneinfo** Directory](#execute-bash-command-\5---create-zoneinfo-directory)
        - [Execute Bash Command \#6 - Download Target Code Tar Archive](#execute-bash-command-\6---download-target-code-tar-archive)
        - [Execute Bash Command \#7 - Download Target Data Tar Archive](#execute-bash-command-\7---download-target-data-tar-archive)
        - [Execute Bash Command \#8 - Unzip Code tar](#execute-bash-command-\8---unzip-code-tar)
        - [Execute Bash Command \#9 - Unzip Data tar](#execute-bash-command-\9---unzip-data-tar)
        - [Execute Bash Command \#10 - Build Time Zone Files](#execute-bash-command-\10---build-time-zone-files)
        - [Execute Bash Command \#11 - Change Directory to zoneinfo](#execute-bash-command-\11---change-directory-to-zoneinfo)
        - [Execute Bash Command \#12 - Delete Old **zoneinfo.zip files**](#execute-bash-command-\12---delete-old-zoneinfozip-files)
        - [Execute Bash Command \#13 - Zip the **zipinfo** Directory](#execute-bash-command-\13---zip-the-zipinfo-directory)
        - [Execute Bash Command \#14 - Change Directory to **D:\\tz**](#execute-bash-command-\14---change-directory-to-d\\tz)
        - [Method-2 Verify the Time Zone Version Contained in zoneinfo.zip](#method-2-verify-the-time-zone-version-contained-in-zoneinfozip)
        - [Method-2 Bash Command Wrap-Up](#method-2-bash-command-wrap-up)
        
      + [Method-2 Step-6 Copy **zoneinfo.zip** to the **Go** Time Directory](#method-2-step-6-copy-zoneinfozip-to-the-go-time-directory)
      
      + [Method-2 Step-7 All Done - Mission Accomplished](#method-2-step-7-all-done---mission-accomplished)      
      
      + [Method-2: Pro’s & Con’s](#method-2-pros--cons)
        - [Method-2 Advantages](#method-2-advantages)
        
        - [Method-2 Drawbacks](#method-2-drawbacks)
        
        - [Method-2 Other Considerations](#method-2-other-considerations)
        
          
    
  + [Clean-Up](#clean-up)
    
    - [Documenting the Current Time Zone Version](#documenting-the-current-time-zone-version)
    
    - [Clean Up Scratch Directory D:\\tz](#clean-up-scratch-directory-d\\tz)
    
      
    
  + [Final Thoughts](#final-thoughts)
    
    - [ianatzformatInfo](#ianatzformatinfo)
    - [datetimeops Package](#datetimeops-package)
    - [Be Careful Out There](#be-careful-out-there)



------



## Introduction

### IANA Time Zone Information

The [**Go Programming Language**](https://golang.org/) incorporates time zone information extracted from the IANA Time Zone database. (IANA stands for Internet Assigned Numbers Authority.)

This database is widely recognized as a leading authority on world time zones.

For some background, reference:

​		<https://en.wikipedia.org/wiki/List_of_tz_database_time_zones>

​		<https://en.wikipedia.org/wiki/Tz_database>

The IANA Time Zone data base information is located at:

​		<https://www.iana.org/time-zones>

​		<https://data.iana.org/time-zones/releases/>



### **Go** Relies on the IANA Time Zone Database

The *Go* Programming Language uses IANA Time Zones in date-time
calculations.

Reference:

​		<https://golang.org/pkg/time/#LoadLocation>


​		<https://golang.org/pkg/time/>



### How Time Zones Are Utilized in **Go** Source Code

#### Load Location

> The **Go** [**Time Package**](https://golang.org/pkg/time/) contains a function named [**LoadLocation**](https://golang.org/pkg/time/#LoadLocation). This function is used to load time zones for date/time calculations. The following was taken from the **Go** Time Package documentation.
>
> **func LoadLocation(name string) (\*Location, error)**
>
> **LoadLocation** returns the Location with the given **name**.

  - > If the **name** is "" or “UTC”, LoadLocation returns [**UTC**](https://en.wikipedia.org/wiki/Coordinated_Universal_Time).
    > If the **name** is "Local", **LoadLocation** returns the location, “Local”. (Local represents the system's local time zone.)

  - > Otherwise, the **name** is taken to be a location name corresponding to a file in the [**IANA Time Zone
    > database**](https://www.iana.org/time-zones), such as "America/New\_York".

  - > The time zone database needed by **LoadLocation** may not be present on all systems, especially non-Unix systems. **LoadLocation** looks in the directory or uncompressed zip file named by the **ZONEINFO** environment variable, if any, then looks in known installation locations on Unix systems,  and finally looks in **$GOROOT/lib/time/zoneinfo.zip**.



#### Time Package References

> The following reference articles provide information on time zones and the **Go** Time Package:
>
>​		[**Go Time Package**](https://golang.org/pkg/time/)
>
>​		[**Detect if ZONEINFO fails in Go - Stack Overflow**](https://stackoverflow.com/questions/51053321/detect-if-zoneinfo-fails-in-go)
>
> ​		[**Package time - Mastering Go**](https://code-examples.net/en/docs/go/time/index)
>
>​		[**knozone**](https://github.com/slotheroo/knozone)



### _**Go**_ Provides for IANA Database Updates

The Time Zone database is updated as warranted by IANA. Usually, these time zone modifications are initiated by nation states and may, or may not, occur intermittently, throughout the year.

Often, new versions of the **Go** Programming Language software will include updated versions of the time zone database. However, there is no guarantee that **Go** will automatically update the time zone data base with each new release.

In order to stay current with time zone modifications, the end-user has the option of updating the local **Go** Programming Language software installation with the latest version of the IANA Time Zone Database.

This update operation must be performed by applying the manual procedures described below. These procedures have been tested using **Go** Version 1.13.1.



## Existing Go Time Zone Configuration

### Configuration Overview

#### Directory/Folder Structure

> The **Go** Programming Language code, compiler and tooling software is typically installed on a drive in the **Go** directory or folder. A common Windows installation would by default, install **Go** on C-Drive as follows:
>
> ​                    				**C:\\Go**
>
> 
>
> On Linux, the **Go** directory is usually found in something like:
>
> ​                               **/usr/local/go**
>
> 
>
> For a discussion of the **Go** directory and related paths,
> reference:
>
> ​		[**Understanding the GOPATH by Gopher Guides**](https://www.digitalocean.com/community/tutorials/understanding-the-gopath)
>
> ​		[**GOPATH, GOROOT, GOBIN - Essential Go**](https://www.programming-books.io/essential/go/gopath-goroot-gobin-d6da4b8481f94757bae43be1fdfa9e73)
>
> ​		[**What is GOROOT and GOPATH and how to fix it?**](https://www.learnsteps.com/goroot-gopath/)
>
> For the purposes of this document, all the following examples assume that **Go** software binaries are installed under **C:\\Go**. If your system is configured for a different **Go** directory, please make the appropriate substitution in the examples shown below.



#### The Go Time Directory

> Time zone information is stored in the **Go** directory structure under the **Go\\lib\\time** sub-directory. Following the Windows example shown above, this would be styled as:

> ​                                   **C:\\Go\\lib\\time**

#### zoneinfo.zip

> All time zone information is contained within a single uncompressed zip file residing in the **Go\\lib\\time** directory named, **zoneinfo.zip**. Using the example above this would equate to a path, file name, file extension of:

> ​	                              **C:\\Go\\lib\\time\\zoneifno.zip**

#### update.bash

> The **Go** installer also creates a bash file in the **Go** time directory named **update.bash**. This file contains a bash script which is used to generate a new **zoneinfo.zip** file. Following the example cited above, the path, file name and file extension for this file would be styled as follows:
>
> ​								**C:\\Go\\lib\\time\\update.bash**



### Upgrading Go Time Zone Configuration

To upgrade the time zone data for your **Go** installation, it will be necessary to create a new version of the **zoneinfo.zip** file containing the desired version of IANA Time Zone information.



## Objective - Create New File: **zoineinfo.zip**

The objective of the time zone update procedure is to create a new uncompressed zip file named **zoneinfo.zip** and copy it to the appropriate **Go** time directory thereby replacing the old version of **zoneinfo.zip**. Once this operation is completed, the new version of **zoneinfo.zip** will be used by **Go** to correctly process time zones anywhere in the world using the specified version of the IANA time
zone database.

After building a new version of the *zoneinfo.zip* file, make sure it is copied to the **Go** time directory:

​									**C:\\Go\\lib\\time\\zoneinfo.zip**



## Two Methods for Updating Go Time Zone Data



### Method-1 The Easy Way: **update.bash**

#### 	Method-1 Step-1 Safety First: Make a Backup Copy of **zoneinfo.zip**

> Make a backup copy of the time zone information file:
>
> ​					**C:\Go\lib\time\zoneinfo.zip**
>
> Store this file some place safe and accessible.



#### 	Method-1 Step-2 Identify the Desired IANA Time Zone Data Version

##### 	Method-1 Access the IANA Time Zone Website

> You have two choices. If you seek the latest version of time zone data, point your browser at:
>
>​					<https://www.iana.org/time-zones>
>
>On the other hand, if you are targeting an earlier version of time zone data, point your browser at the release history web site:
>
> ​					[https://data.iana.org/time-zones/releases/](https://data.iana.org/time-zones/releases/)

##### 	Method-1 Record the Target Time Zone Version Number

> Make a note of the desired time zone version. Example: **2019c**.



#### Method-1 Step-3 Create a Scratch Directory

Create a new, empty scratch directory which can be used for time zone update operations. This is a temporary directory which will be used for intermediate processing and file storage. After the update procedure is completed, you should be able to delete this directory. Example:

> ​							**D:\\tz**



#### Method-1 Step-4 Copy *update.bash* to Scratch Directory

The bash script file **update.bash** resides in the **Go** time directory:
						**C:\\Go\\lib\\time\\update.bash**

Copy **update.bash** to the new empty, scratch directory created in **Step-3**, above.

> ​		**Copy -\>**
>
> ​					**From:**      **C:\\Go\\lib\\time\\update.bash**
>
> ​					**To:**            **D:\\tz\\update.bash**



#### Method-1 Step-5 Modify the Bash File

After copying the file **update.bash** to the scratch directory, it will be necessary to modify this bash script using a text editor.

1.  Open the scratch directory version of **update.bash** in a text editor (Reference **Step-4** above).

2.  Within the **update.bash** file, locate the bash script section labeled **“\# Versions to use.”** in your text editor.

3.  Modify the “**CODE**” and “**DATA**” variables to reflect the desired time zone version captured in **Step-2 b**, above.

> ​     				**From:**
>
> ​          						**CODE=2019b**
>
> ​          						**DATA=2019b**
>
> ​     			 	**To:**
>
> ​          						**CODE=2019c**
>
> ​          						**DATA=2019c**
>
> * Instead of **“2019c”** you will of course use the desired IANA time zone version captured in **Step-2 b**, above.

4.  Save the modified **update.bash** file and exit your text editor.



#### Method-1 Step-6 Navigate to Scratch Directory

1.  Open a Terminal Window. On Windows, this involves opening a [Cygwin](https://www.cygwin.com/) terminal window or a terminal
    emulator configured for Cygwin, such as [ConEmu](https://conemu.github.io/).

2.  Change directories to the scratch directory - i.e., make the scratch directory the current directory.

> ​	On Windows this means using **Cygwin**. Change Directory to scratch directory, **D:\\tz**.
>
>​									**cd /cygdrive/d**
>
>​									**cd ./tz**
>
>To verify the current working directory, at the command line, execute the command **pwd**. This will list and verify the current working directory.



#### Method-1 Step-7 Execute the Bash Script

1.  IMPORTANT: Before starting this step, make certain that **Step-5**, above, has been completed correctly.

2.  Now, execute the bash script.

> ​					**./update.bash**
>
> Be advised, this command will generate a lot of messages in the terminal window.

3.  If all went well and no errors were returned, the following final message should be displayed in the terminal window:

> ​					**“New time zone files in zoneinfo.zip.”**



#### Method-1 Step-8 Close the Terminal Windows

Close the Terminal Window.



#### Method-1 Step-9 Examine Scratch Directory Contents

> Assuming a successful bash script execution, only two files should remain in the scratch directory, **D:\\tz**:

1.  **update.bash**

2.  **zoneinfo.zip**

> This **zoneinfo.zip** file is the new replacement zip file configured with the target IANA Time Zone information selected in **Step-2 b**, above.



#### Method-1 Step-10 Copy **zoneinfo.zip** to the **Go** Time Directory

> Copy the newly created **zoneinfo.zip** file from the scratch directory to the **Go** time directory.

1.  After ensuring you have backup copy (See **Step-1**, above), delete the old **zoneinfo.zip** file from the **Go** time directory.

> ​					**Delete File:** **C:\\Go\\lib\\time\\zoneinfo.zip**

2.  Copy the new **zoneinfo.zip** file from the scratch directory to the **Go** time directory.

> ​		**Copy:**
>
> ​			**From:**	**D:\\tz\\zoneinfo.zip**
>
> ​			**To:**		  **C:\\Go\\lib\\time\\zoneinfo.zip**



#### Method-1 Step-11 All Done - Mission Accomplished

> That’s it. The objective has been achieved. By copying the file **zoneinfo.zip** to the **Go** time directory **(C:\\Go\\lib\\time)**, updated time zone information is now configured. Henceforth, the **Go** compiler will incorporate this updated time zone information into all date/time calculations executed in code.
>
> Be sure to read the section entitled **Clean-Up** below.



#### Method-1: Pro’s and Con’s

##### Method-1 Advantages

> **Method-1** is easy. Compared to **Method-2**, creation of the file **zoneinfo.zip** file requires far fewer manual operations and therefore incurs less chance of error. My reading indicates that the **Go** team recommends use of **Method-1**.



##### Method-1 Drawbacks

> Unfortunately, there is no version information contained within **zoneinfo.zip**. Therefore, you will not be  able to confirm the IANA version number of time zone data contained in **zoneinfo.zip** from an examination of that file. However, if the **CODE** and **DATA** variables were updated correctly as outlined in **Step-5**, above, there is a high probability that the **update.bash** script executed correctly and that the new **zoneinfo.zip** file contains the correct version time zone data.
> 
>Still, there is no way to verify success through empirical observation. Operational success is totally dependent on the bash script.



### Method-2: The Not So Easy Way

This method involves downloading the IANA Time Zone files and manually building the **zoneinfo.zip** file.

#### Method-2 Step-1 Safety First: Make a Backup Copy of **zoneinfo.zip**

Make a backup copy of the time zone information file:

​					**C:\\Go\\lib\\time\\zoneinfo.zip.**

> Store this file some place safe and accessible.



#### Method-2 Step-2 Identify the Desired IANA Time Zone Data Version

##### Method-2 Access the IANA Time Zone Website

> You have two choices. If you seek the latest version of time zone data, point your browser at:
>
> ​			[https://www.iana.org/time-zones](https://www.iana.org/time-zones)
>
>On the other hand, if you are targeting an earlier version of time zone data, point your browser at the release history web site:
>
> ​			[https://data.iana.org/time-zones/releases/](https://data.iana.org/time-zones/releases/)



##### Method-2 Record the Target Time Zone Version Number

> Make a note of the desired time zone version. Example: **2019c**.



#### Method-2 Step-3 Create a Scratch Directory

Create a new, empty scratch directory which can be used for time zone update operations. This is a temporary directory which will be used for intermediate processing and file storage. After the update procedure is completed, you should be able to delete this directory. Example:

> ​					**D:\\tz**



#### Method-2 Step-4 Navigate to Scratch Directory

1.  Open a Terminal Window. On Windows, this involves opening a [Cygwin](https://www.cygwin.com/) terminal window or a terminal
    emulator configured for Cygwin, such as [ConEmu](https://conemu.github.io/).

2.  Change directories to the scratch directory - i.e., make the scratch directory the current directory.

> On Windows this means using **Cygwin**. Change Directory to scratch directory, **D:\\tz**.
>
>​					**cd /cygdrive/d**
>
>​					**cd ./tz**
>
>To verify the current working directory, at the command line, execute the command **pwd**. This will list and verify the current working directory.



#### Method-2 Step-5 Execute Bash Commands

> Note: The following bash commands were identified by analyzing and dissecting the bash file **C:\\Go\\lib\\time\\update.bash** identified in the section on **Method-1**, above.
>
> After making the scratch directory (**D:\\tz**) the current directory, set the environment and create the sub-directories by issuing the following series of bash commands. Again, you must be in the scratch directory (**D:\\tz**) when issuing these commands.



##### Execute Bash Command \#1 - Set Environment to Exit on Error

> ​					**set -e**
>
> This sets the environment such that scripts will exit on error.



##### Execute Bash Command \#2 - Delete Old **work** directory tree

> ​					**rm -rf work**
>
> This deletes the sub-directory **work** and the child directories in
> that tree.



##### Execute Bash Command \#3 - Create **work** Directory

> ​					**mkdir ./work**
>
> This creates the sub-directory **work** (**D:/tz/work**). To verify, execute command **“ls -lh”** to list contents of the **D:/tz** directory.



##### Execute Bash Command \#4 - Change to **work** Directory

> ​					**cd ./work**
>
> Change directory to sub-directory **work (D:/tz/work)**. Execute the command **pwd** to list and verify the current working directory.



##### Execute Bash Command \#5 - Create **zoneinfo** Directory

> ​					**mkdir ./zoneinfo**
>
> This command creates a new sub-directory, **zoneinfo (D:\\tz\\work\\zoneinfo)**. To verify, execute the command **“ls  -lh”** and list the contents of the **D:\\tz\\work** directory.



##### Execute Bash Command \#6 - Download Target Code Tar Archive

> Replace the text phrase **2019c** in the command syntax below with the target time zone version identified in **Step-2 b** above.
>
>**curl -L -O https://www.iana.org/time-zones/repository/releases/tzcode2019c.tar.gz**
>
> This **curl** command must be executed on one line from the **work** directory **(D:\\tz\\work)**.  To verify download of the correct file, execute the “**ls -lh**” command to list the contents of the **work** directory **(D:\\tz\\work)**.



##### Execute Bash Command \#7 - Download Target Data Tar Archive

> Replace the text phrase **2019c** in the command syntax below with the target time zone version identified in **Step-2,b** above.
>
>**curl -L -O https://www.iana.org/time-zones/repository/releases/tzdata2019c.tar.gz**
>
> This **curl** command must be executed on one line from the **work** directory (**D:\\tz\\work**). To verify download of the correct file, execute the “**ls -lh**” command to list the contents of the **work** directory **(D:\\tz\\work)**.



##### Execute Bash Command \#8 - Unzip Code tar

> Replace the text phrase **2019c** in the command syntax below with the latest time zone version acquired   in **Step-2** above. This command will unzip the file downloaded with **Bash Command \# 6**, above.
>
> ​					**tar xzf tzcode2019c.tar.gz**
>
>This command must be executed on one line from the **work** directory **(D:\\tz\\work)**.



##### Execute Bash Command \#9 - Unzip Data tar

> Replace the text phrase **2019c** in the command syntax below with the latest time zone version acquired in **Step-2** above. This command will unzip the file downloaded with **Bash Command \# 7**, above.
>
> ​					**tar xzf tzdata2019c.tar.gz**
>
>This command must be executed on one line from the **work** directory **(D:\\tz\\work)**.



##### Execute Bash Command \#10 - Build Time Zone Files

> This will execute the **make** command and create the actual time zone data files.
>
>​	**make CFLAGS=-DSTD\_INSPIRED AWK=awk TZDIR=zoneinfo posix\_only**
>
>This command must be executed on one line from the **work** directory **(D:\\tz\\work)**. After completion, the **zoneinfo** sub-directory **(D:\\tz\\work\\zoneinfo)** should be populated with time zone files.



##### Execute Bash Command \#11 - Change Directory to zoneinfo

> Change directory to **D:\\tz\\work\\zoneinfo**
>
> ​					**cd ./zoneinfo**
>
> If uncertain about your current directory, run bash command, **pwd** to list and verify the current working directory.
>
>It is important to be positioned in the correct directory before continuing.



##### Execute Bash Command \#12 - Delete Old **zoneinfo.zip files**

> As a precaution, this command is run from **D:\\tz\\work\\zoneinfo**. It will effectively delete old **zoneinfo.zip** files which may reside in directory **D:\\tz**, two directories above the current directory. The new **zipinfo.zip** file will be created in **D:\\tz**.
>
> ​					**rm -f ../../zoneinfo.zip**



##### Execute Bash Command \#13 - Zip the **zipinfo** Directory

> This zip command will create the **zoneinfo.zip** file containing the specified time zone information downloaded in **Bash Commands \# 6 and 7**, above.
>
> ​					**zip -0 -r ../../zoneinfo.zip**
>
> **“-0”** (Zero) option = “store files with no compression”



##### Execute Bash Command \#14 - Change Directory to **D:\\tz**

> ​					**cd ../..**
>
> This assumes that the change directory command is being executed from directory, **D:\\tz\\work\\zoneinfo**. After executing this “**change directory**” command, the current working directory should be **D:\\tz**.
>
> To list and verify the current working directory, issue the **pwd** command.



##### Method-2 Verify the Time Zone Version Contained in zoneinfo.zip

> You now have the option of verifying to a certainty that the newly created file, **D:\\tz\\zoneinfo.zip**, contains the correct version of time zone data.
>
> Open your text editor of choice and load the file **D:\\tz\\work\\version** (“**version**” has no file extension).
>This file was downloaded under **Bash Command \#7** described above.  The **version** file will contain a version number associated with the downloaded time zone data.  This version number should match the target time zone data version selected in **Step-2 b**, above.



##### Method-2 Bash Command Wrap-Up

> Close the terminal window.
>
> This concludes the bash commands required to create the new **zoneinfo.zip** file containing the time zone information downloaded in **Bash Commands \#’s 6 and 7**.  If all steps were completed successfully and no errors were encountered, the new **zipinfo.zip** file resides in the base scratch directory, **D:\\tz**.
>
> The following steps involve copying this **zipinfo.zip** file to the **Go** time directory **C:\\Go\\lib\\time)**.  After completing this operation, consider taking action to clean up the scratch directory, **D:\\tz**.



#### Method-2 Step-6 Copy **zoneinfo.zip** to the **Go** Time Directory

Copy the newly created **zoneinfo.zip** file from the scratch directory to the **Go** time directory.

1.  After ensuring you have backup copy (See **Step-1**, above), delete the old **zoneinfo.zip** file from the **Go** time directory.

>
>
> ​					**Delete File: C:\\Go\\lib\\time\\zoneinfo.zip**
>
>

2.  Copy the new **zoneinfo.zip** file from the scratch directory to the **Go** time directory.

> ​			**Copy:**
>
> ​				**From:**		**D:\\tz\\zoneinfo.zip**
>
> ​				**To:** 			**C:\\Go\\lib\\time\\zoneinfo.zip**



#### Method-2 Step-7 All Done - Mission Accomplished

> That’s it. The objective has been achieved. By copying the file **zoneifno.zip** to the **Go** time directory **(C:\\Go\\lib\\time)**, updated time zone information is now properly configured. Henceforth, the **Go** compiler will incorporate this updated time zone information into all date/time calculations executed in code.
>
> Be sure to read the section entitled **Clean-Up** below.



#### Method-2: Pro’s & Con’s

##### Method-2 Advantages

> **Method-2** provides more granular control of the time zone upgrade process.  This increased flexibility allows the end user to specify any version of the time zone database in the IANA Time Zone release history.  More importantly, this method allows for verification of the version number associated with the time zone data downloaded and incorporated into the final **zoneinfo.zip** file.  This verification capability is superior to that provided in **Method-1**.



##### Method-2 Drawbacks

> Compared to **Method-1**, **Method-2** has many keyboard intensive, manual operations which must be performed by the end-user.  The chance of error is therefore increased.



##### Method-2 Other Considerations

> As mentioned in section **Method-1: Pro’s & Con’s**, **Method-1**, which relies on the bash script **C:\\Go\\lib\\time\\update.bash,** is the recommended approach for updating time zones.   The keyboard intensive bash commands used in **Method-2** were extracted from **C:\\Go\\lib\\time\\update.bash**.  If you use **Method-2**, it would be prudent to examine and review the bash commands as new versions of the Go Programming Language are released.



## Clean-Up

### Documenting the Current Time Zone Version

After completing the update procedure, give some thought to documenting the version of time zone data now residing in the **Go** time directory **(C:\\Go\\lib\\time)**.

Each time the **C:\\Go\\lib\\time** directory is updated with a new version of **Go** software downloaded from [**golang.org**](https://golang.org/), two files in the **Go** time directory are **overwritten**:

> ​					**C:\\Go\\lib\\time\\zoneinfo.zip**
>
> ​									And
>
> ​					**C:\\Go\\lib\\time\\update.bash**

In this case, the only source of time zone version numbers is found through an examination of the **CODE** and **DATA** variables contained in **update.bash** file.

Therefore, if you choose to manually update time zone data using one of the two methods outlined above, you may want to include a custom text file in the **Go** time directory **(C:\\Go\\lib\\time)** documenting the actual version of time zone data contained in **zoneinfo.zip**.  Remember, the time zone data version cannot be verified through an examination of the **zoneinfo.zip** file.

If a custom text file is created documenting the time zone data version, be aware that it may become outdated immediately after installing a newer version of **Go** Programming Language software from [**golang.org**](https://golang.org/).



### Clean Up Scratch Directory D:\\tz

Remember that both **Method-1** and **Method-2** required the creation of an empty scratch directory. The example given was a directory named, **D:\\tz**.  This directory now contains file and depending on future usage plans, you may want to delete it.





## Final Thoughts

### ianatzformatInfo

The application **ianatzformatInfo** may prove helpful.  This application is located on GitHub at [https://github.com/MikeAustin71/ianatzformatInfo](https://github.com/MikeAustin71/ianatzformatInfo).

It is designed to read raw IANA Time Zone data files and create a **Go** source file containing types and associated data structures which may be used to facilitate access to over 680 IANA and Military time zone across the globe. Using the same time zone data download technique outlined in **Method-2** above, the end-user can easily update time zone data as needed.

### datetimeops Package

The **datetimeops** package provides many data types and helper methods designed to manage the full gamut of date time calculations.  It is located on GitHub at [https://github.com/MikeAustin71/datetimeopsgo](https://github.com/MikeAustin71/datetimeopsgo).

Among its many features, the **datetimeops** packages gives the developer full access to the time zone data structures generated by the **ianatzformatInfo** application.



### Be Careful Out There

Managing time zone information and updates involves complexity and the ever-present risk of error.  Weigh the update option carefully in light of your application’s specific requirements.  

Not all time zones updates are significant.  The IANA Time Zone database provides an optional email subscription which will detail the nature of time zone changes.  Your final decision on time zone upgrades should take into account the risk/return trade-offs associated time zone changes which have little or no effect on your operations. 
