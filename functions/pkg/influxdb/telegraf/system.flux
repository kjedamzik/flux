package telegraf

import "influxdb"

option bucket = "telegraf"

cpu = (bucket=bucket) =>
    influxdb.measurement(bucket:bucket,m:"cpu")
