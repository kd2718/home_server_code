#include <Arduino.h>

const int powerPin = 25; // Powers the sensor
const int adcPin = 33;   // Reads the data
const int ledPin = 2;   // Reads the data

// Calibration Data (From your bench test)
const int DryValue = 2370;
const int WetValue = 400;

void setup()
{
    Serial.begin(115200);

    // Set up the power pin and ensure it starts OFF
    pinMode(powerPin, OUTPUT);
    pinMode(ledPin, OUTPUT);
    digitalWrite(powerPin, LOW);
    delay(500);
}

void loop()
{
    // --- 1. Power ON ---
    //digitalWrite(powerPin, HIGH);
    digitalWrite(powerPin, HIGH);
    digitalWrite(ledPin, HIGH);

    // --- 2. Stabilize ---
    // Capacitive sensors need ~50-100ms to charge the electric field
    delay(5000);

    // --- 3. Read ---
    uint16_t rawValue = analogRead(adcPin);
    rawValue = analogRead(adcPin);

    // --- 4. Power OFF ---
    digitalWrite(powerPin, LOW);
    digitalWrite(ledPin, LOW);
    delay(5000);
    uint16_t rawValueLow = analogRead(adcPin);

    // --- 5. Convert to Percentage ---
    // map(value, fromLow, fromHigh, toLow, toHigh)
    // Note: We swap the order (Dry, Wet) because the sensor works backwards (Low = Wet)
    int percent = map(rawValue, DryValue, WetValue, 0, 100);

    // Clamp the values to keep it 0-100%
    if (percent < 0)
        percent = 0;
    if (percent > 100)
        percent = 100;

    Serial.printf("Raw HIGH: %d | Moisture: %d%%\n", rawValue, percent);
    Serial.printf("Raw LOW: %d | Moisture: %d%%\n", rawValueLow, percent);

}
/*
#include <Arduino.h>

const int powerPin = 25; // Make sure your wire is in Pin 25!
const int adcPin = 33;   // Make sure signal is in Pin 33!

void setup()
{
    Serial.begin(115200);
    pinMode(powerPin, OUTPUT);

    // Turn ON and keep ON
    digitalWrite(powerPin, LOW);

    // Wait 1 full second for the sensor to wake up
    delay(1000);
}

void loop()
{
    int rawValue = analogRead(adcPin);
    Serial.printf("Sensor Reading: %d\n", rawValue);
    delay(1000);
}
*/
