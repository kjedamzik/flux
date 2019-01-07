package influxdb

measurement = (bucket,m) =>
    from(bucket:bucket)
        |> filter(fn:(r) => r._measurement == m)
