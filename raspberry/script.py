#!/usr/bin/env python

import time
import grovepi
from subprocess import call as system
# import urllib.request
import urllib
import urllib2
from os.path import isfile


if not isfile("/tmp/already_start") :
    system(["touch", "/tmp/already_start"])
    system(["pulseaudio", "--start"])
    time.sleep(10)
user = "Jean-Yves"
LED = 2 
SW = 4
numleds = 1 #If you only plug 1 LED, change 10 to 1

grovepi.pinMode(LED,"OUTPUT")
grovepi.pinMode(SW,"INPUT")
time.sleep(1)

testColorBlack = 0   # 0b000 #000000
testColorBlue = 1    # 0b001 #0000FF
testColorGreen = 2   # 0b010 #00FF00
testColorCyan = 3    # 0b011 #00FFFF
testColorRed = 4     # 0b100 #FF0000
testColorMagenta = 5 # 0b101 #FF00FF
testColorYellow = 6  # 0b110 #FFFF00
testColorWhite = 7   # 0b111 #FFFFFF

def turn_off_led(led_num):
    grovepi.storeColor(0,0,0)
    time.sleep(.1)
    grovepi.chainableRgbLed_pattern(LED, 0, led_num)
    time.sleep(.1)
  
#Turns on the led specified by led_num to color set by r,g,b
def turn_on_led(led_num,r,g,b):
    grovepi.storeColor(r,g,b)
    time.sleep(.1)
    grovepi.chainableRgbLed_pattern(LED, 0, led_num)
    time.sleep(.1)
    
try:

    print("Initialise")
    message="Initialisation de la clef d'or." 
    system(["./dit.sh", message])
    # init chain of leds
    grovepi.chainableRgbLed_init(LED, numleds)
    time.sleep(.5)
    turn_on_led(0,0,255,0)
    time.sleep(1)
    turn_off_led(0)
    #why did we need to do that
    # grovepi.pinMode(SW,"INPUT")
    while 1 :
        #grovepi.pinMode(SW,"INPUT")
        print(grovepi.digitalRead(SW))
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

        time.sleep(1)
except KeyboardInterrupt:
    # reset (all off)
    grovepi.chainableRgbLed_test(LED, numleds, testColorBlack)
except IOError:
    print ("Error")
