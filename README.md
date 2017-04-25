[![Build Status](https://travis-ci.org/rvelhote/timestamp-marshal.svg?branch=master)](https://travis-ci.org/rvelhote/timestamp-marshal) [![codecov](https://codecov.io/gh/rvelhote/timestamp-marshal/branch/master/graph/badge.svg)](https://codecov.io/gh/rvelhote/timestamp-marshal) [![Code Climate](https://codeclimate.com/github/rvelhote/timestamp-marshal/badges/gpa.svg)](https://codeclimate.com/github/rvelhote/timestamp-marshal) [![Issue Count](https://codeclimate.com/github/rvelhote/timestamp-marshal/badges/issue_count.svg)](https://codeclimate.com/github/rvelhote/timestamp-marshal)

# Unix Timestamp Marshal/Unmarshal
This package will allow you to Marshal and Unmarshal Unix Timestamps from and to JSON datasets. This is useful when 
you have an API that returns or accepts unix timestamps as a value which the time package is not able to deal with.

It will create a new data type named `timestamp.Unix` that extends/implements `time.Time` so after the marshaling 
business is done you can use it like a normal `time.Time` data type.

Marshaling will convert the `time.Time` field into its unix timestamp equivalent.
Unmarshaling will convert a unix timestamp into its `time.Time` equivalent.

Please be aware that, because timestamps do not contain timeszones, you will have to deal with that yourself.

# Usage
You can make use of this package by defining the field(s) of your structs as `timestamp.Unix` and mapping the field 
to the field name of your JSON data structure.

```
type PeerInfo struct {
    Id int `json:"id"`
    Addr string `json:"addr"`
    AddrLocal string `json:"addrlocal"`
    LastSend timestamp.Unix `json:"lastsend"`
    LastRecv timestamp.Unix `json:"lastrecv"`
}
```

The just use it as you'd use a `time.Time` field. Add, subtract, format... you know the deal.
