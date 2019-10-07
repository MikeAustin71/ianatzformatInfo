# ianatzformatInfo

This program is written in the *Go* programming language, a.k.a. *golang*.
The application executable is, *ianatzformatInfo.exe*.

The objective of this application is to parse and format time zone information
for use in *Go* programs. 

## Table of Contents
+ [Source Of Time Zone Data](#source-of-time-zone-data)
+ [Application Processing](#application-processing)
+ [Base Input Data](#base-input-data)
    - [Input File Example](#input-file-example)
        - [Line 1](#line-1)
        - [Line 2](#line-2)
        - [Configuration](#configuration)
+ [IANA Time Zone Information](#iana-time-zone-information)
    - [Latest Time Zone Database](#latest-time-zone-database)
    - [Previous Time Zone Databases](#previous-time-zone-databases)
+ [Application Output](#application-output)
    - [timezonedata.go](#timezonedatago)
+ [Questions and Comments](#questions-and-comments)

## Source Of Time Zone Data

The IANA Time Zone database is widely recognized as the the world's
leading authority on time zones.

    Reference:
    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
    https://en.wikipedia.org/wiki/Tz_database
    https://www.iana.org/time-zones.

The *Go* Programming Language uses IANA Time Zones in date-time
calculations.

    Reference:
    https://golang.org/pkg/time/#LoadLocation

## Application Processing
Using this application, information is extracted from IANA Time Zone
source files and reformatted as a series of 'type' data structures
stored in a source file suitable for use in a *Go* program.

This application will process IANA source data files and create a series
of data structures and types housed in a source file used by the 'Go'
programming language. These types will facilitate the use of global time
zones in date-time operations incorporated in *Go* programs. 

## Base Input Data 
In order to function properly, this application expects to find the text file
[targettzdata.txt](./app/input/targettzdata.txt) in a sub-directory labeled *input* which is located directly
beneath the directory which houses this application executable. 

#### Input File Example
Assume application executable, *ianatzformatInfo.exe*,
is located in 'D:\myAppDir' .
     
Input File [targettzdata.txt](./app/input/targettzdata.txt) MUST BE located in directory: 

        'D:\myAppDir\input\targettzdata.txt'.

[targettzdata.txt](./app/input/targettzdata.txt) is a text file containing two lines
of information on the first two lines in the text file.
    
Each line of text must be terminated with a new line
character, '\n'. 

#####Line 1 
    The first line designates the 'path' to the IANA time zone data files.
    The application will search this path for valid IANA time zone database
    files which will be used as input.

#####Line 2
    The second line designates the 'path' to the output file.
    
    The application will create the output file 'timezonedata.go'
    in this output directory.  This 'Go' source file will contain
    data structures used to access world wide time zones. 

##### Configuration
Configure these two lines in accordance with the following example.

    Example:
    "InputDirectory: D:\T11\data\tzdata2019c\n"
    "OutputDirectory: D:\GoProjects\ianatzformatInfo\app\output\n"

The leading field names, 'InputDirectory:' and 'OutputDirectory:'
are mandatory.


## IANA Time Zone Information

Reference the IANA Time Zone Databases.

#### Latest Time Zone Database

    http://www.iana.org/time-zones

#### Previous Time Zone Databases

    https://data.iana.org/time-zones/releases/

## Application Output

The IANA time zone database contains text files which represents
raw data on world wide time zones. This application is designed
to parse and format this information before generating a *Go*
source file containing a series of types defining world time
zones. 

#### timezonedata.go
The output source file, [timezonedata.go](./app/output/timezonedata.go),
is created in the output directory specified in the input file,
'targettzdata.txt'.


## Questions And Comments

    mike.go@paladinacs.net
        


