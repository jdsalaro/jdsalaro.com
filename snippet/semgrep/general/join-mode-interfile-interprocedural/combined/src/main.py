from collections import namedtuple

from utils import SINK

# [!] alternatively:
# import utils
# [!] either way use fully qualified imports

Source = namedtuple("Source", "data")
RAW = Source({"mykey": "myval"})

# ---------------------------------------------


def simple_one_param(source=None):
    myval = source.data["mykey"]
    tainted = "Hello '{}' World!".format(myval)

    # ruleid: match
    return SINK.sink(tainted)


# ---------------------------------------------


def simple_multi_param(source=None, other=None):
    myval = other.data["mykey"]
    tainted = "Hello '{}' World!".format(myval)

    # ruleid: match
    return SINK.sink(tainted)


# ---------------------------------------------


def interprocedural_source_helper(source=None):
    myval = source.data["mykey"]
    tainted = "Hello '{}' World!".format(myval)

    return tainted


def interprocedural_source():
    # ruleid: match
    return SINK.sink(interprocedural_source_helper(source))


# ---------------------------------------------


def interprocedural_no_source_helper():
    myval = RAW.data["mykey"]
    tainted = "Hello '{}' World!".format(myval)
    return tainted


def interprocedural_no_source(source=None):
    # ruleid: match
    return SINK.sink(interprocedural_no_source_helper())


# ---------------------------------------------

from utils import SOURCE


def interprocedural_interfile(source=None):
    # ruleid: match
    return SINK.sink(SOURCE.interfile_source(source))


def interprocedural_interfile_no_source(source=None):
    # ruleid: match
    return SINK.sink(SOURCE.interfile_source())


# ---------------------------------------------

if __name__ == "__main__":
    simple_one_param(RAW)
    simple_multi_param(RAW, RAW)
    interprocedural_source()
    interprocedural_no_source()
    interprocedural_interfile(RAW)
    interprocedural_interfile_no_source(RAW)
