inData = "
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,dateTime:RFC3339,double,string,string,string,string,string,string
#group,false,false,true,true,false,false,true,true,true,true,true,true
#default,_result,,,,,,,,,,,
,result,table,_start,_stop,_time,_value,_field,_measurement,device,fstype,host,path
,,96,2018-05-22T19:53:26Z,2018-05-22T19:54:16Z,2018-05-22T19:53:26Z,34.98234271799806,used_percent,disk,disk1s1,apfs,host.local,/
,,96,2018-05-22T19:53:26Z,2018-05-22T19:54:16Z,2018-05-22T19:53:36Z,34.98234941084654,used_percent,disk,disk1s1,apfs,host.local,/
,,96,2018-05-22T19:53:26Z,2018-05-22T19:54:16Z,2018-05-22T19:53:46Z,34.982447293755506,used_percent,disk,disk1s1,apfs,host.local,/
,,96,2018-05-22T19:53:26Z,2018-05-22T19:54:16Z,2018-05-22T19:53:56Z,34.982447293755506,used_percent,disk,disk1s1,apfs,host.remote,/
,,96,2018-05-22T19:53:26Z,2018-05-22T19:54:16Z,2018-05-22T19:54:06Z,34.98204153981662,used_percent,disk,disk1s1,apfs,host.remote,/
,,96,2018-05-22T19:53:26Z,2018-05-22T19:54:16Z,2018-05-22T19:54:16Z,34.982252364543626,used_percent,disk,disk1s1,apfs,host.remote,/
"
outData = "
#datatype,string,long,dateTime:RFC3339,string,string,double,double
#group,false,false,true,true,true,false,false
#default,0,,,,,,
,result,table,_stop,_field,_measurement,host.local,host.remote
,,0,2018-05-22T19:54:16Z,used_percent,disk,34.98237980753337,34.98224706603858
"

t_pivot_mean = (table=<-) =>
  table
  |> range(start: 2018-05-22T19:53:26Z)
  |> group(columns: ["_stop", "_measurement", "_field", "host"])
  |> mean()
  |> pivot(rowKey: ["_stop"], columnKey: ["host"], valueColumn: "_value")
  |> yield(name:"0")

testingTest(name: "pivot_mean",
            input: testLoadStorage(csv: inData),
            want: testLoadMem(csv: outData),
            test: t_pivot_mean)