# Description:


An example in Go for publishing binary-encoded data to a socket and consuming the published data using two different goroutines.

The program sends data to a socket in a binary-encoded format. The data simulates sensor readings and has the following structure:

```plaintext
0       1       2       3       4       5       6       7
0123456701234567012345670123456701234567012345670123456701234567
+-------+-------+-------+-------+-------+-------+-------+------
|          SensorID             |    LocationID |     Timestamp
+-------+-------+-------+-------+-------+-------+-------+------
                |      Temp     |
+-------+-------+-------+-------+
```

The listener consumes the data from the socket, decodes it, and prints it to the screen.

### TODO: make publisher and listerner two different entry points
 So that they can be started in two different containers and create docker compose file to simulate behaviour