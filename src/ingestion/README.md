# adm-tsmanager-service

TSManager service handles timescale database containing data sent from devices.
It allows to make custom queries retrieving data filtering on different fields.

### HTTP-GET-example

An example of HTTP GET:
http://api.zerinth.com/v1/tsmanager/workspace/{workspaceid}/tag/{tag}?deviceID={devID}&start={start}&end={end}&custom={custom}

device_id: string
start, end: date YYYY-MM-DD hh:mm (RFC3339)

custom: custom query conditions with '|' character as delimiter
```
/workspace/example/tag/caffee?deviceID=dev123&limit=500&start=2019-12-31 15:30&end=2019-12-31 17:00&custom=|field1=0|field3>food|field4|
```

When custom is set to a value like:
```
	"custom=|field=1|" or "custom=|field|"
```
It will show ONLY 'field' of payload.

To get all field when custom is not null, nedd all n-fields like:
```
	custom=|field1>5|field2|...|fieldn|
```
In the '...' have to write explicitly the fields needed.
