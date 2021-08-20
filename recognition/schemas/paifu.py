from pydantic import BaseModel


class PaifuResponse(BaseModel):
    hands: list[str]
    bakaze: str
    jikaze: str
    dora: list[str]
    junme: int
