
from fastapi import FastAPI
from pydantic import BaseModel
import random


app = FastAPI()
class InputData(BaseModel):
    text: str

@app.post("/predict")
def predict(data: InputData):
    sentiment = random.choice(["positive", "negative", "neutral"])
    return {"sentiment": sentiment}
   

