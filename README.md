# addressbook
Written by Kirby Flake (kirbyef@gmail.com)

This is an address book server written in golang that uses REST API.

The only external API required is Gorilla web toolkit. Please look at 
http://www.gorillatoolkit.org/pkg/mux how to add it to your GOPATH.

Each record contains the following data:

Go Name | JSON Name | Description
:---: | :---: | :---:
ID | id | Id number 
FirstName | firstname | First Name
LastName | lastname | Last Name
Email | email | Email Address
Phone | phone | Phone Number

This API uses standard REST semantics. To list, add, modify or delete an entry a users must make a request using the appropriate method:

Method | URL | Description
:----: | :---: | :---:
GET | /records | Lists all records
GET | /record/{id} | Lists record with ID specified
POST | /record/{id} | Adds new record with ID specified 
DELETE | /record/{id} | Deletes record with ID specified
PATCH | /record/{id} | Modifies record with ID specified

If an Id is not found, it will return a JSON response with "" data for each field.

An example json file is included examplerecord.json to show the format of the JSON data for adding or modifing a record. 

All the data is stored on the server in a file called addressbook.csv. This allows easy importing and exporting of the data. To import an addressbook, just upload a new addressbook.csv to the server. To export the addressbook, just download the /addressbook.csv file.  

