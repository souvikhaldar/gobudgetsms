# Gobudgetsms

This library is used to easily make use of the Budget SMS public API to send sms, plus it lets to create message of variable length. 
## Getting Started
Just insert the details according to the example provided below and get going. No dependencies are requied to use this lib, very lightweight.

Demo working sample of this library can be found at https://youtu.be/D62KMuHezZg


![alt text](https://github.com/souvikhaldar/gobudgetsms/blob/master/Screen%20Shot%202018-07-17%20at%208.56.13%20PM.png)


### Prerequisites

You need to have an account on Budget SMS.

### Installing

```
go get github.com/souvikhaldar/gobudgetsms
```

## Example 

Details regarding account can be found here https://www.budgetsms.net/controlpanel/api-details/

Fill the parameters according to https://www.budgetsms.net/sms-http-api/test-sms/
```
// Just one function call to send the sms
res , err := gobudgetsms.SendSMS(detail Details)
if err != nil {
	//handle error
}

// An example of detail object
detail := Details{
	"username",	
	"userid",	
	"handle",	
	"The message",
	"from",
	"to",
	"",
	0,
	0,
	0,
}
```



## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
