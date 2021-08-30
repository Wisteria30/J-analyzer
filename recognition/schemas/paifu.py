from pydantic import BaseModel


class PaifuResponse(BaseModel):
    zikaze: int
    bakaze: int
    turn: int
    dora_indicators: list[int]
    hand_tiles: list[int]
