#include <Arduino.h>

// GPIO 2 is the standard onboard LED for most ESP32 DevKitC boards
const int ledPin = 2;
const int a2dPin=33;
const int poutPin=4;


void setup()
{
    // Initialize the digital pin as an output
    pinMode(ledPin, OUTPUT);
    Serial.begin(115200);
    pinMode(poutPin, OUTPUT);
    digitalWrite(poutPin,HIGH);
    digitalWrite(ledPin, HIGH);
}

void loop()
{
    digitalWrite(ledPin, LOW);
    uint16_t sensorValue = analogRead(a2dPin);
    delay(1000);
    Serial.println("Test 123");
    Serial.printf("sensor reading: %d\n", sensorValue);

}
