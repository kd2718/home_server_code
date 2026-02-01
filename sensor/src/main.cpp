#include <Arduino.h>

// GPIO 2 is the standard onboard LED for most ESP32 DevKitC boards
const int ledPin = 2;
const int a2dPin=33;
const int poutPin=23;


void setup()
{
    // set bod rate
    Serial.begin(115200);

    // Initialize the digital pin as an output
    pinMode(ledPin, OUTPUT);
    pinMode(poutPin, OUTPUT);

    // set initial pin output
    digitalWrite(poutPin,HIGH);
    digitalWrite(ledPin, HIGH);

}

void loop()
{
    digitalWrite(poutPin,HIGH);
    digitalWrite(ledPin, HIGH);
    uint16_t sensorValue = analogRead(a2dPin);
    Serial.printf("sensor reading HIGH: %d\n", sensorValue);
    delay(5000);

    digitalWrite(poutPin,LOW);
    digitalWrite(ledPin, LOW);
    sensorValue = analogRead(a2dPin);
    Serial.printf("sensor reading LOW: %d\n", sensorValue);
    delay(5000);
}
