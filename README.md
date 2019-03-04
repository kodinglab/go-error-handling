# go-error-handling
Go error handling and custom logger with gin-gionic

Middleware for error handling and logger which display Controller, Method or Function name and also show line number where the Error was raised.

Best fit for RESTful API 

## Explanation
Look at [middleware/logger.go](https://github.com/kodinglab/go-error-handling/blob/master/middleware/logger.go) there are three methods:
* Error - To show error message from application
* Info - Write custom log info, *example* for debugging
* ErrorOutput - Display output error in JSON format (500 status code)

Error files will be formed by date and located in **logs** directory
```
logs/logs-2019-03-04.log
logs/logs-2019-03-05.log
```

## Usage
Write log error in somewhere location (Controller, Model or anything you want)
```go
....
....

type TestController struct {
}

func (ctr *TestController) GetErrorRoute(c *gin.Context) {
    // Raising error
    err := errors.New("This is error route")
    if err != nil {
        //.... your logic here....
    
        // write error log
        middleware.Log{}.Error(err)
    
        // and you can display output error (500 status code)
        middleware.ErrorOutput(c, err)
    }
}

```
Log output
```bash
[ERROR] 2019/03/05 - 03:12:47 | error-handling/controllers.(*TestController).GetErrorRoute[/home/alfa/apps/go/src/error-handling/controllers/test.go:27]
This is error route
```
Log format
```bash
[ErrorType] Timestamp | Controller - Method[path/to/file:LineNumber]
```
Display output
```json
{
    "message": "This is error route",
    "status": false
}
```

## Authors
**Alfa** - *Initial work* - [kodinglab](https://github.com/kodinglab)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
