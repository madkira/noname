import serial
import time

rpiser = serial.Serial('/dev/ttyAMA0', baudrate=9600, timeout=0)
rpiser.flush()

while rpiser.inWaiting() == 0:
    time.sleep(0.005)

serial_data = rpiser.readline()
print(serial_data)
