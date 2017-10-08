# DalhousieScheduleBuilderAPI
API for the dal schedule builder web app. Being used in my dal schedule builder which can be found at http://dalschedulebuilder.com

## Usage
Simply build and run the api.go and it will start the REST API.
The API takes in 2 parameters, the subject code (i.e CSCI) and the term (1 = fall, 2 = winter).

dalschedulebuilder.com/api/courses?s=`SUBJECT CODE`&t=`TERM`

## Example Response
dalschedulebuilder.com/api/courses?s=CSCI&t=1
Returns all the CSCI courses for fall term.

Some courses have been removed for simplicity sake.
```json
{
  "data": [
    {
      "category": "CSCI",
      "code": "1108",
      "title": "Experimental Robotics",
      "classes": [
        {
          "id": 726,
          "crn": 14739,
          "section": "01",
          "type": "Lec",
          "credit_hours": 3,
          "days": "MON, WED",
          "times": "0935-1025",
          "location": "Studley LSC-PSYCHOLOGY P5260",
          "max": "90",
          "current": "74",
          "waitlist": "NA",
          "prof": "Trappenberg T."
        },
        {
          "id": 727,
          "crn": 14740,
          "section": "B01",
          "type": "Lab",
          "credit_hours": 0,
          "days": "MON, WED",
          "times": "1135-1325",
          "location": "Studley COMPUTER SCIENCE 134",
          "max": "30",
          "current": "30",
          "waitlist": "NA",
          "prof": "Staff"
        },
        {
          "id": 728,
          "crn": 14741,
          "section": "B02",
          "type": "Lab",
          "credit_hours": 0,
          "days": "MON, WED",
          "times": "1135-1325",
          "location": "Studley COMPUTER SCIENCE 228",
          "max": "20",
          "current": "8",
          "waitlist": "NA",
          "prof": "Staff"
        },
        {
          "id": 729,
          "crn": 14742,
          "section": "B03",
          "type": "Lab",
          "credit_hours": 0,
          "days": "MON, WED",
          "times": "1635-1825",
          "location": "Studley COMPUTER SCIENCE 134",
          "max": "30",
          "current": "30",
          "waitlist": "NA",
          "prof": "Staff"
        }
      ]
    }
  ]
}
```

## TODO
* Add pages to API request
