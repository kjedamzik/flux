package testdata

inData = "
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,string,string,long
#group,false,false,true,true,false,true,true,false
#default,,,,,,,,
,result,table,_start,_stop,_time,_measurement,_field,_value
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:30Z,_m,FF,1
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:40Z,_m,FF,1
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,FF,1
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:00Z,_m,FF,1
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:10Z,_m,FF,1
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:20Z,_m,FF,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:30Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:40Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:00Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:10Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:30Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:40Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:00Z,_m,QQ,1
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:10Z,_m,RR,1
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:30Z,_m,RR,1
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:40Z,_m,RR,1
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,RR,1
,,3,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:00Z,_m,SR,1
,,3,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:10Z,_m,SR,1
,,3,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:20Z,_m,SR,1
"
outData = "
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,string,string,long
#group,false,false,true,true,false,true,true,false
#default,_result,,,,,,,
,result,table,_start,_stop,_time,_measurement,_field,_value
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:30Z,_m,FF,1
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:40Z,_m,FF,2
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,FF,3
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:00Z,_m,FF,4
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:10Z,_m,FF,5
,,0,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:20Z,_m,FF,6
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:30Z,_m,QQ,1
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:40Z,_m,QQ,2
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,QQ,3
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:00Z,_m,QQ,4
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:10Z,_m,QQ,5
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:30Z,_m,QQ,6
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:40Z,_m,QQ,7
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,QQ,8
,,1,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:00Z,_m,QQ,9
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:10Z,_m,RR,1
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:30Z,_m,RR,2
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:40Z,_m,RR,3
,,2,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:53:50Z,_m,RR,4
,,3,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:00Z,_m,SR,1
,,3,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:10Z,_m,SR,2
,,3,2018-05-22T19:53:00Z,2018-05-22T19:54:50Z,2018-05-22T19:54:20Z,_m,SR,3
"

t_cumulative_sum_default = (table=<-) => table
  |> cumulativeSum()

testingTest(
    name: "cumulative_sum_default",
    input: testLoadStorage(csv: inData),
            want: testLoadMem(csv: outData),
    test: t_cumulative_sum_default
)
