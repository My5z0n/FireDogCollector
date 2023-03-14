import os
import subprocess
import json
import yaml

RESULT_FILENAME = r""".\result.json"""
BASE = r""".\batch_base.yaml"""
BATCH = r""".\batch.yaml"""
MACROBASE_COMAND = r""" java -jar .\macrobase-legacy.jar pipeline .\batch.yaml"""


class CallMacroBase():

    def prepareBatch(self, macrobase_attributes, base_query):
        with open(BASE) as f:
            data = yaml.load(f, Loader=yaml.SafeLoader)
            data["macrobase.loader.db.baseQuery"] = macrobase_attributes
            data["macrobase.loader.attributes"] = base_query
            with open(BATCH, "w") as f:
                yaml.dump(data, f)
        print(data)

    def makeCall(self):

        if os.path.exists(RESULT_FILENAME):
            os.remove(RESULT_FILENAME)

        process = subprocess.Popen(
            MACROBASE_COMAND, stdout=subprocess.PIPE, shell=True)
        process.communicate()

        result = ""
        job = None
        if os.path.exists(RESULT_FILENAME):
            with open(RESULT_FILENAME, "r") as json_file:
                json_object = json.load(json_file)
                job = json_object
                result = json.dumps(json_object)
                print(json.dumps(json_object, indent=1))

        else:
            print("No result file")

        return job


if __name__ == "__main__":
    tmp = CallMacroBase()
    tmp.makeCall()
