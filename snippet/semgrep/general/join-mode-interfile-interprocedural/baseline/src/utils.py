from collections import namedtuple


class SINK:
    def sink(data=None):
        if "myval" in data:
            print("VULNERABLE!")
            return False
        else:
            print("SAFE!")
            return True


class SOURCE:
    Source = namedtuple("Source", "data")
    RAW = Source({"mykey": "myval"})

    def interfile_source(source=None):
        if source is None:
            source = SOURCE.RAW

        myval = source.data["mykey"]
        tainted = "Hello '{}' World!".format(myval)

        return tainted
