from Services.CallMacroBase import CallMacroBase
from Services.ClickhouseRepository import ClickhouseRepository


class AnomalyService():
    client: any

    def __init__(self) -> None:
        self.macro_base = CallMacroBase()
        self.db = ClickhouseRepository()

    def find_outlines(self, goodPart: str, badPart: str):

        x2 = goodPart  # '!START#/api/user/:user'
        x1 = badPart  # '!START#/api/user/:user#!END'
        l1 = goodPart.split("#")  # ['!START', '/api/user/:user']
        l1 = ["'"+l+"'" for l in l1]
        l3 = ",".join(l1)
        baseSQLStatement = f"SELECT startsWith(t.paths, '{x1}') AS Test, s.* FROM spans s, traces t WHERE s.trace_id = t.trace_id AND startsWith(t.paths, '{x2}') AND s.span_name IN ({l3});"
        colNames = self.db.getColumnNames("spans")
        self.macro_base.prepareBatch(baseSQLStatement, colNames)
        result = self.macro_base.makeCall()

        return result


if __name__ == "__main__":
    x = AnomalyService()
    x.find_outlines("!START#/api/user/:user", "!START#/api/user/:user#!END",)
