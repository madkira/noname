#!/usr/bin/env python

import time
import subprocess # import call as system
import paho.mqtt.client as paho

# import urllib.request
from os.path import isfile

brocker="localhost"

def on_message(client, userdata, message):
    time.sleep(1)
    subprocess.run(["./dit.sh",str(message.payload.decode("utf-8"))])

client= paho.Client("client-001")
client.on_message=on_message
client.connect(brocker)
client.loop_start()

client.subscribe("goldenkey/exit")
print("Initialise")
message="Initialisation de la clef d'or." 
subprocess.run(["./dit.sh", message])
# init chain of leds
#why did we need to do that
try :
    while 1 :
        '''
        result = subprocess.run(['./desfirefuncs',''], stdout=subprocess.PIPE).stdout.decode('utf-8')
        if(result) :
            #print(result)
            if(current =="") :
                current = users[result]
                subprocess.Popen(["mosquitto_pub -h localhost -t goldenkey/entry -m \"%s\"" % current], shell = True)
                print(current)
        else :
            if(current!="") :
                subprocess.Popen(["mosquitto_pub -h localhost -t goldenkey/exit -m \"%s\"" % current], shell = True)
                current = ""
        if grovepi.digitalRead(SW) :
            message =' '.join(["Bon retour parmi nous", user, "; vous"])
            url = 'http://192.168.1.190/whoishere/add'
            data = urllib.urlencode({'id' : '1'})
            req = urllib2.Request(url=url,data=data)
            nbmessage = urllib2.urlopen(req).read()
            if nbmessage :
                nbmessage = urllib.urlopen("http://192.168.1.190/mail/1").read()
                if nbmessage :
                    message=' '.join([message, "avez",str(nbmessage),"messages non lut"])
                else : 
                    message=' '.join([message, "n'avez pas de nouveaux messages"])
                system(["./dit.sh", message])
            while grovepi.digitalRead(SW) :
                time.sleep(1)
            message="Bonne journee"+user
            url ='http://192.168.1.190/whoishere/rm'
            data = urllib.urlencode({'id' : '1'})
            req = urllib2.Request(url=url,data=data)
            urllib2.urlopen(req).read()
            system(["./dit.sh", message])
'''
except KeyboardInterrupt:
    # reset (all off)
    print ("key error")
    #grovepi.chainableRgbLed_test(LED, numleds, testColorBlack)
except IOError:
    print ("Error")


client.disconnect()
client.loop_stop()
