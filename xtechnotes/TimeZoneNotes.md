Updating Go Installations
with the
Latest IANA Time Zone Data

## Table of Contents

[Introduction 1](#introduction)

[IANA Time Zone Information 1](#iana-time-zone-information)

[Go Relies on IANA Time Zone Database
1](#go-relies-on-the-iana-time-zone-database)

[Go Provides for IANA Database Updates
2](#go-provides-for-iana-database-updates)

[Existing Go Time Zone Configuration
3](#existing-go-time-zone-configuration)

[Configuration Overview 3](#configuration-overview)

[Directory/Folder Structure 3](#directoryfolder-structure)

[The Go Time Directory 3](#_Toc21648258)

[zoneinfo.zip 3](#zoneinfo.zip)

[update.bash 4](#update.bash)

[Upgrading Go Time Zone Configuration
4](#upgrading-go-time-zone-configuration)

[Objective - Create New File: *zoineinfo.zip*
5](#objective---create-new-file-zoineinfo.zip)

[Two Methods for Updating Go Time Zone Data
6](#two-methods-for-updating-go-time-zone-data)

[Method-1 The Easy Way: ***update.bash***
6](#method-1-the-easy-way-update.bash)

[Step-1 “Safety First” 6](#step-1-safety-first)

[Step-2 Copy *update.bash* to Scratch Directory *D:\\tz*
6](#step-2-copy-update.bash-to-scratch-directory-dtz)

[Step-3 Modify *update.bash* 7](#step-3-modify-update.bash)

[Step-4 Execute *update.bash* Script
8](#step-4-execute-update.bash-script)

[Step-5 Copy zoneinfo.zip to the Go Time Directory
9](#step-5-copy-zoneinfo.zip-to-the-go-time-directory)

[Method-1: Pro’s and Con’s 10](#method-1-pros-and-cons)

[Advantages 10](#advantages)

[Drawbacks 10](#drawbacks)

[Method-2: The Not So Easy Way 11](#method-2-the-not-so-easy-way)

[Step-1 “Safety First” 11](#step-1-safety-first-1)

[Step-2 Identify the Latest IANA Time Zone Data Version
11](#step-2-identify-the-latest-iana-time-zone-data-version)

[Access the IANA Time Zone Website
11](#access-the-iana-time-zone-website)

[Capture the Latest Time Zone Version Number
11](#capture-the-latest-time-zone-version-number)

[Step-3 Create a Scratch Directory
11](#step-3-create-a-scratch-directory)

[Step-3 Navigate to Scratch Directory
12](#step-3-navigate-to-scratch-directory)

[Step-4 Execute Bash Commands 12](#step-4-execute-bash-commands)

[Execute Bash Command \#1 - Set Environment to Exit on Error
12](#execute-bash-command-1---set-environment-to-exit-on-error)

[Execute Bash Command \#2 - Delete Old ***work*** directory tree
12](#execute-bash-command-2---delete-old-work-directory-tree)

[Execute Bash Command \#3 - Create ***work*** Directory
13](#execute-bash-command-3---create-work-directory)

[Execute Bash Command \#4 - Change to ***work*** Directory
13](#execute-bash-command-4---change-to-work-directory)

[Execute Bash Command \#5 - Create ***zoneinfo*** Directory
13](#execute-bash-command-5---create-zoneinfo-directory)

[Execute Bash Command \#6 - Download Latest Code
13](#execute-bash-command-6---download-latest-code)

[Execute Bash Command \#7 - Download Latest Data
14](#execute-bash-command-7---download-latest-data)

[Execute Bash Command \#8 - Unzip Code tar
14](#execute-bash-command-8---unzip-code-tar)

[Execute Bash Command \#9 - Unzip Data tar
14](#execute-bash-command-9---unzip-data-tar)

[Execute Bash Command \#10 - Build Time Zone Files
15](#execute-bash-command-10---build-time-zone-files)

[Execute Bash Command \#11 - Change Directory to zoneinfo
15](#execute-bash-command-11---change-directory-to-zoneinfo)

[Execute Bash Command \#12 - Delete Old **zoneinfo.zip files**
15](#execute-bash-command-12---delete-old-zoneinfo.zip-files)

[Execute Bash Command \#13 - Zip the **zipinfo** Directory
16](#execute-bash-command-13---zip-the-zipinfo-directory)

[Execute Bash Command \#14 - Change Directory to **D:\\tz**
16](#execute-bash-command-14---change-directory-to-dtz)

[Verify the Time Zone Version Contained in zoneinfo.zip
16](#verify-the-time-zone-version-contained-in-zoneinfo.zip)

[Bash Command Wrap-Up 17](#bash-command-wrap-up)

[Step-5 Copy zoneinfo.zip to the Go Time Directory
17](#step-5-copy-zoneinfo.zip-to-the-go-time-directory-1)

[Method-2: Pro’s & Con’s 18](#method-2-pros-cons)

[Advantages 18](#advantages-1)

[Drawbacks 18](#drawbacks-1)

[Clean-Up 19](#clean-up)

[How Time Zones Are Utilized in ***Go*** Source Code
20](#how-time-zones-are-utilized-in-go-source-code)

[Load Location 20](#load-location)

## Introduction

### IANA Time Zone Information

The ***Go*** Programming Language incorporates time zone information
extracted from the IANA Time Zone database. This database is widely
recognized as a leading authority on world time zones.

Reference:

<https://en.wikipedia.org/wiki/List_of_tz_database_time_zones>

<https://en.wikipedia.org/wiki/Tz_database>

The IANA Time Zone data base and reference information is located at:

<https://www.iana.org/time-zones>

<https://data.iana.org/time-zones/releases/>

### ***Go*** Relies on the IANA Time Zone Database

The *Go* Programming Language uses IANA Time Zones in date-time
calculations.

Reference:

<https://golang.org/pkg/time/#LoadLocation>

<https://golang.org/pkg/time/>

### ***Go*** Provides for IANA Database Updates

The Time Zone database is updated as warranted by the Internet Assigned
Numbers Authority (IANA). Time zone modifications are usually initiated
by nation states and may or may not occur intermittently, throughout the
year. In order to stay current with time zone modifications, the
end-user has the option of updating the local ***Go*** Programming
Language software installation with the latest version of the IANA Time
Zone Database. ***Go*** does not automatically update the time zone data
base. That task is therefore left to the end-user to be performed
through the manual update procedure described below.

## Existing Go Time Zone Configuration

### Configuration Overview

#### Directory/Folder Structure

> The ***Go*** Programming Language binaries are typically installed on
> a drive in the ***Go*** directory or folder. A common Windows
> installation would install ***Go*** on C-Drive as follows:

***C:\\Go***

> <span id="_Toc21648258" class="anchor"></span>Your system may vary
> such that the ***Go*** binaries could be installed under a different
> directory.
>
> For the purposes of this document, all the following examples assume
> that the ***Go*** software is installed under ***C:\\Go***. If your
> system is configured for a different ***Go*** directory, please make
> the appropriate substitution in the examples shown below.

#### The Go Time Directory

> Time zone information is stored in the ***Go*** directory structure
> under the ***Go\\lib\\time*** sub-directory. Following the example
> above, this would be styled as:

**C:\\Go\\lib\\time**

#### zoneinfo.zip

> All time zone information is contained within a single uncompressed
> zip file residing in the ***Go\\lib\\time*** directory named,
> ***zoneinfo.zip***. Using the example above this would equate to a
> path, file name, file extension of:

***C:\\Go\\lib\\time\\zoneifno.zip***

#### update.bash

> The ***Go*** installer also creates a bash file in the Go\\lib\\time
> directory named ***update.bash***. This bash file is used to generate
> a new ***zoneinfo.zip*** file. Following the example cited above the
> path, file name and file extension for this file would be styled as
> follows:

***C:\\Go\\lib\\time\\update.bash***

### Upgrading Go Time Zone Configuration

To upgrade the time zone data for your ***Go*** installation, it will be
necessary to create a new version of the ***zoneinfo.zip*** file
containing the desired version of IANA Time Zone information.

## Objective - Create New File: ***zoineinfo.zip***

The objective of the time zone update procedure is to create an
uncompressed zip file named *zoneinfo.zip* and copy it to the
appropriate directory thereby replacing the old version of
***zoneinfo.zip***. Once this operation is completed, the new version of
***zoneinfo.zip*** will be used by ***Go*** to correctly process time
zones anywhere in the world using the specified version of the IANA time
zone database.

When you finish building the latest version of the *zoneinfo.zip* file,
make sure it is copied to the ***Go*** time directory:

***C:\\Go\\lib\\time\\zoneinfo.zip***

## Two Methods for Updating Go Time Zone Data

### Method-1 The Easy Way: ***update.bash***

#### Step-1 “Safety First”

1.  Make a backup copy of the time zone information file
    ***$GOROOT/lib/time/zoneinfo.zip***.

2.  Create a new, empty scratch directory which can be used for time
    zone update operations. Example: ***D:\\tz***.

#### Step-2 Copy *update.bash* to Scratch Directory *D:\\tz*

1.  The File ***update.bash*** resides in the Go time directory:
    ***$GOROOT/lib/time/update.bash***

2.  Copy ***update.bash*** to the new empty, scratch directory created
    in Step-1, above.

> **Copy -\>**
>
> From: ***$GOROOT/lib/time/update.bash***
>
> To: ***D:\\tz\\update.bash***

#### Step-3 Modify *update.bash*

> After copying update.bash to the scratch directory, it will be
> necessary to modify this bash script using a text editor.

1.  Point your browser at <https://www.iana.org/time-zones> .

2.  Make a note of the latest time zone version. Example: ***2019c***

3.  Open the scratch directory version of ***update.bash*** in a text
    editor (see Step-2 above: ***D:\\tz\\update.bash***).

4.  Locate the bash script section labeled **“\# Versions to use.”** in
    your text editor.

5.  Modify the “CODE” and “DATA” variables to reflect the latest time
    zone version captured in Step-3 a, above.

> **From:**
>
> **CODE=2019b**
>
> **DATA=2019b**
>
> **To:**
>
> **CODE=2019c**
>
> **DATA=2019c**
>
> Instead of **“2019c”** you will of course use the latest IANA time
> zone version.

6.  Save the modified ***update.bash*** file and exit your text editor.

#### Step-4 Execute *update.bash* Script

1.  Open a terminal window and change directories to the scratch
    directory (i.e. make the scratch directory the current window.

> On windows this means using ***Cygwin***. Change Directory to scratch
> directory, ***D:\\tz***.
>
> **cd /cygdrive/d**
>
> **cd ./tz**

2.  Now, execute the bash script.

> **./update.bash**

3.  If all went well and no errors were emitted, the following final
    message should be displayed in the terminal window:

> **“New time zone files in zoneinfo.zip.”**

4.  Close the Terminal Window

5.  Assuming a successful bash script execution, only two files should
    remain in the scratch directory:

<!-- end list -->

1.  **update.bash**

2.  **zoneinfo.zip**

> This ***zoneinfo.zip*** file is the new replacement file configured
> with the latest IANA Time Zone information.

#### Step-5 Copy zoneinfo.zip to the Go Time Directory

> Copy the newly created zoneinfo.zip file from the scratch directory to
> the Go time directory.

1.  After ensuring you have backup copy, delete the old zoneinfo.zip
    file from the Go Time Directory.

> Delete File: ***$GOROOT/lib/time/zoneinfo.zip***

2.  Copy the new zoneinfo.zip file from the scratch directory to the Go
    Time directory.

> Copy:
>
> From: ***D:\\tz\\zoneinfo.zip***
>
> To: **$GOROOT/lib/time/zoneinfo.zip**

3.  Done - Mission Accomplished

> That’s it. The objective has been achieved. By copying the file
> **zoneifno.zip** to the Go Time directory **($GOROOT/lib/time)**,
> updated time zone information is now configured. The Go compiler will
> now incorporate this updated time zone information into all date/time
> calculations executed in code.

#### Method-1: Pro’s and Con’s

##### Advantages

> Method-1 is easy. Compared to Method-2, creation of the file
> ***zoneinfo.zip*** file requires far fewer manual operations.

##### Drawbacks

> Unfortunately, there is no version information contained within
> ***zoneinfo.zip***. Therefore, you will not be able to confirm the
> IANA version number of time zone data contained in ***zoneinfo.zip***
> from an examination of that file. However, if the ***CODE*** and
> ***DATA*** variables were updated correctly as outlined in **Step-3**
> above, there is a high probability that the **update.bash** script
> executed correctly and that the new ***zoneinfo.zip*** file contains
> the correct time zone data.

### Method-2: The Not So Easy Way

This method involves downloading the IANA Time Zone files and manually
building the **zoneinfo.zip** file.

#### Step-1 “Safety First”

Make a backup copy of the time zone information file:

> ***$GOROOT/lib/time/zoneinfo.zip.***

#### Step-2 Identify the Latest IANA Time Zone Data Version

##### Access the IANA Time Zone Website

> Point your browser at <https://www.iana.org/time-zones>

##### Capture the Latest Time Zone Version Number

> Make a note of the latest time zone version. Example: **2019c**.

#### Step-3 Create a Scratch Directory

Create a new, empty scratch directory which can be used for time zone
update operations. Example:

> ***D:\\tz.***

#### Step-3 Navigate to Scratch Directory

1.  Open a Terminal Window.

2.  Make the scratch directory the current directory. On windows this
    means using Cygwin. Change Directory to scratch directory,
    ***D:\\tz***.

> **cd /cygdrive/d**
>
> **cd ./tz**

#### Step-4 Execute Bash Commands

> After making the scratch directory (***D:\\tz***) the current
> directory, set the environment and create the sub-directories by
> issuing the following commands. Again, you must be in the scratch
> directory (***D:\\tz***) when issuing these bash commands.

##### Execute Bash Command \#1 - Set Environment to Exit on Error

> **set -e**
>
> This sets the environment such that scripts will exit on error.

##### Execute Bash Command \#2 - Delete Old ***work*** directory tree

> **rm -rf work**
>
> This deletes the sub-directory ***work*** and child directories in the
> tree.

##### Execute Bash Command \#3 - Create ***work*** Directory

> **mkdir ./work**
>
> This creates the sub-directory ***work*** (***D:/tz/work***). To
> verify, execute command “***ls***” to list contents of the ***D:/tz***
> directory.

##### Execute Bash Command \#4 - Change to ***work*** Directory

> **cd ./work**
>
> Change directory to sub-directory **work (D:/tz/work)**. To verify,
> execute command “***pwd***” to confirm the current working directory.

##### Execute Bash Command \#5 - Create ***zoneinfo*** Directory

> **mkdir ./zoneinfo**
>
> Creates new sub-directory, ***zoneinfo (D:\\tz\\work\\zoneinfo)***. To
> verify, execute the command “***ls***” and list the contents of the
> ***D:\\tz\\work*** directory.

##### Execute Bash Command \#6 - Download Latest Code

> Replace the text phrase ***2019c*** in the command syntax below with
> the latest time zone version acquired in Step-2 above.
>
> **curl -L -O
> https://www.iana.org/time-zones/repository/releases/tzcode2019c.tar.gz**
>
> This **curl** command must be executed on one line from the **work**
> directory ***(D:\\tz\\work)***. To verify download of the correct
> file, execute the “***ls -lh***” command to list the contents of the
> work directory ***(D:\\tz\\work)***.

##### Execute Bash Command \#7 - Download Latest Data

> Replace the text phrase ***2019c*** in the command syntax below with
> the latest time zone version acquired in Step-2 above.
>
> **curl -L -O
> https://www.iana.org/time-zones/repository/releases/tzdata2019c.tar.gz**
>
> This **curl** command must be executed on one line from the **work**
> directory (***D:\\tz\\work***). To verify download of the correct
> file, execute the “***ls -lh***” command to list the contents of the
> work directory ***(D:\\tz\\work)***.

##### Execute Bash Command \#8 - Unzip Code tar

> Replace the text phrase ***2019c*** in the command syntax below with
> the latest time zone version acquired in Step-2 above. This command
> will unzip the file downloaded with Command \# 6, above.
>
> **tar xzf tzcode2019c.tar.gz**
>
> This command must be executed on one line from the **work** directory
> **(D:\\tz\\work)**.

##### Execute Bash Command \#9 - Unzip Data tar

> Replace the text phrase ***2019c*** in the command syntax below with
> the latest time zone version acquired in Step-2 above. This command
> will unzip the file downloaded with Command \# 7, above.
>
> **tar xzf tzdata2019c.tar.gz**
>
> This command must be executed on one line from the **work** directory
> **(D:\\tz\\work)**.

##### Execute Bash Command \#10 - Build Time Zone Files

> This will execute the ***make*** command and create the actual time
> zone data files.
>
> **make CFLAGS=-DSTD\_INSPIRED AWK=awk TZDIR=zoneinfo posix\_only**
>
> This command must be executed on one line from the **work** directory
> **(D:\\tz\\work)**. The ***zoneinfo*** sub-directory
> **(D:\\tz\\work\\zoneinfo)** should now be populated with time zone
> files.

##### Execute Bash Command \#11 - Change Directory to zoneinfo

> Change directory to ***D:\\tz\\work\\zoneinfo***
>
> **cd ./zoneinfo**
>
> If uncertain about your current directory, run bash command,
> ***pwd***. This will display the current directory.
>
> It is important to be positioned in the correct directory.

##### Execute Bash Command \#12 - Delete Old **zoneinfo.zip files**

> As a precaution, this command is run from
> ***D:\\tz\\work\\zoneinfo***. It will effectively delete old
> ***zoneinfo.zip*** files which reside in directory ***D:\\tz***, two
> directories above the current directory. The new ***zipinfo.zip***
> file will be created in ***D:\\tz***.
>
> **rm -f ../../zoneinfo.zip**

##### Execute Bash Command \#13 - Zip the **zipinfo** Directory

> **zip -0 -r ../../zoneinfo.zip \***
>
> **“-0”** (Zero) option = “store files with no compression”
>
> This zip command will create the ***zoneinfo.zip*** file containing
> the specified time zone information downloaded in Commands \# 6 and 7,
> above.

##### Execute Bash Command \#14 - Change Directory to **D:\\tz**

> **cd ../..**
>
> This assumes that the change directory command is being executed from
> directory, **D:\\tz\\work\\zoneinfo**. After executing this “**change
> directory**” command, the current working directory should be
> ***D:\\tz***. If in doubt about the current directory, issue the
> ***pwd*** command to list the current working directory.

##### Verify the Time Zone Version Contained in zoneinfo.zip

> You now have the option of verifying to a certainty that the newly
> created file, ***D:\\tz\\zoneinfo.zip***, contains the correct version
> of time zone data.
>
> Open your text editor of choice and load the file
> ***D:\\tz\\work\\version*** (“version” has no file extension). This
> file was downloaded under Bash Command \#7 discussed above. The
> ***version*** file will contain a version number associated with the
> downloaded time zone data.

##### Bash Command Wrap-Up

> Close the terminal window.
>
> This concludes the bash commands required to create the new
> ***zoneinfo.zip*** file containing the time zone information
> downloaded in **Commands \# 6 and 7**. If all steps were completed
> successfully and no errors were encountered, the new ***zipinfo.zip***
> file resides in the base scratch directory, ***D:\\tz***.
>
> The following steps involve copying this ***zipinfo.zip*** file to the
> Go time directory **($GOROOT/lib/time)**. After completing this
> operation, consider taking action to clean up scratch directory
> ***D:\\tz***.

#### Step-5 Copy zoneinfo.zip to the Go Time Directory

> Copy the newly created ***zoneinfo.zip*** file from the scratch
> directory to the Go time directory.

1.  After ensuring you have backup copy, delete the old zoneinfo.zip
    file from the Go Time Directory.

> Delete File: **$GOROOT/lib/time/zoneinfo.zip**

2.  Copy the new zoneinfo.zip file from the scratch directory to the Go
    Time directory.

> **Copy:**
>
> From: **D:\\tz\\zoneinfo.zip**
>
> To: **$GOROOT/lib/time/zoneinfo.zip**

**
**

#### Method-2: Pro’s & Con’s

##### Advantages

> Method-2 provides more granular control of the time zone upgrade
> process. This increased flexibility allows the end user to specify any
> version of the time zone database in the release history. More
> importantly, this method allows for verification of the version number
> associated with the time zone data downloaded and incorporated into
> the final ***zoneinfo.zip*** file. This verification capability is
> superior to that provided in Method-1.

##### Drawbacks

> Compared to Method-1, Method-2 has many more moving parts, thereby
> increasing the chance of error. Plus, Method-2 is considerably more
> labor intensive.

## Clean-Up

After completing the update procedure, give some thought to documenting
the version of time zone data now residing in the ***Go*** Time
directory ***($GOROOT/lib/time)***.

Each time the ***$GOROOT*** directory is updated with a new version of
***Go*** software downloaded from ***golang.org***, two files in the
***Go*** time directory are overwritten:
***$GOROOT/lib/time/zoneinfo.zip*** and
***$GOROOT/lib/time/update.bash***.

In this case, the only source of time zone version numbers is found
through an examination of the **CODE** and **DATA** variables contained
in ***update.bash*** file.

Therefore, if you choose to manually update time zone data using one of
the two methods outlined above, you may want to include a text file in
the ***Go*** time directory ***($GOROOT/lib/time)*** documenting the
actual version of time zone data contained in ***zoneinfo.zip***.
Remember, time data version cannot be verified by examining the
zoneinfo.zip file.

Again, be wary of updates the ***Go*** language software installation as
any custom text file information will become quickly outdated.

## How Time Zones Are Utilized in ***Go*** Source Code

### Load Location

The ***Go*** Time Package contains a function named
[***<span class="underline">LoadLocation</span>***](https://golang.org/pkg/time/#LoadLocation).
This function is used to load time zones for date/time calculations.

func LoadLocation(name string) (\*Location, error)

***LoadLocation*** returns the Location with the given name.

If the name is "" or “UTC”, LoadLocation returns ***UTC***. If the name
is "Local", LoadLocation returns ***Local*** (See Go Time Package method
[Local](https://golang.org/pkg/time/#Time.Local)).

Otherwise, the name is taken to be a location name corresponding to a
file in the [*IANA Time Zone
database*](https://www.iana.org/time-zones), such as
"America/New\_York".

The time zone database needed by LoadLocation may not be present on all
systems, especially non-Unix systems. LoadLocation looks in the
directory or uncompressed zip file named by the ***ZONEINFO***
environment variable, if any, then looks in known installation locations
on Unix systems, and finally looks in
***$GOROOT/lib/time/zoneinfo.zip***.
