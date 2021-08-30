from pydantic import BaseModel


class PaifuTestResponse(BaseModel):
    hands: list[str]
    bakaze: str
    jikaze: str
    dora: list[str]
    junme: int
