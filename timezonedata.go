package main




// timeZoneArray - This array contains time zones from the IANA database. 
// The total number of time zones is 10
var timeZoneAry = []string {
           "Asia/Thimphu",
           "Asia/Brunei",
           "Asia/Yangon",
           "Asia/Kolkata",
           "Asia/Almaty",
           "Asia/Oral",
           "Indian/Maldives",
           "Asia/Qatar",
           "Asia/Damascus",
           "Asia/Ashgabat"}


// mapTzLinks - A listing of deprecated time zones with links to active 
// IANA time zones. key='deprecated time zone' value='current active time zone'
// The number of links is: 1
var linkMap = map[string]string {
   "Link":         "Asia/Dubai", 
    }


