import random
from fastapi import FastAPI, Query, Path
from typing import Optional

app = FastAPI()

def resolve_location_and_sensorID(location: Optional[str], sensorID: Optional[str]):
    # If no location is provided, use a default based on sensor ID
    if not location:
        if sensorID == "1":
            location = "Living Room"
        elif sensorID == "2":
            location = "Bedroom"
        elif sensorID == "3":
            location = "Kitchen"
        else:
            location = "Unknown"

    # If no sensor ID is provided, generate one based on location
    if not sensorID:
        if location == "Living Room":
            sensorID = "1"
        elif location == "Bedroom":
            sensorID = "2"
        elif location == "Kitchen":
            sensorID = "3"
        else:
            sensorID = "0"
    return location, sensorID

@app.get("/temperature")
def get_temperature(
    location: Optional[str] = Query(None)
):
    location, sensorID = resolve_location_and_sensorID(location, None)
    return {
        "Location": location,
        "SensorID": sensorID,
        "Value": round(random.uniform(-91.2, 57.8), 1)
    }

@app.get("/temperature/{sensorID}")
def get_temperature_by_path(
    sensorID: str = Path(...),
    location: Optional[str] = Query(None)
):
    location, sensorID = resolve_location_and_sensorID(location, sensorID)
    return {
        "Location": location,
        "SensorID": sensorID,
        "Value": round(random.uniform(-91.2, 57.8), 1)
    } 
