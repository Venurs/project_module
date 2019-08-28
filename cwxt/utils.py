import hashlib
import base64
from typing import Any, Union
from decimal import Decimal
import time
import datetime
from sqlalchemy.orm import class_mapper


def serialize(data, many=False):
    """Transforms a model into a dictionary which can be dumped to JSON."""
    # first we get the names of all the columns on your model
    if many:
        results = list()
        for model in data:
            single_model = dict((m, getattr(model, m)) for m in [c.key for c in class_mapper(model.__class__).columns])
            results.append(single_model)
        return results
    else:
        columns = [c.key for c in class_mapper(data.__class__).columns]
        # then we return their values in a dict
        return dict((c, getattr(data, c)) for c in columns)


def str2md5(original_str):

    md_str = hashlib.md5()
    md_str.update(original_str.encode(encoding='utf-8'))
    md_str = md_str.hexdigest()

    md_str = base64.b64encode(bytes(md_str, encoding='utf-8'))

    new_md_str = hashlib.md5()
    new_md_str.update(str(md_str, 'utf-8').encode(encoding='utf-8'))
    new_md_str = new_md_str.hexdigest()

    return new_md_str


def date2stamp(date_time):
    """
    :param date_time: datetime
    :return: int timestamp
    """
    time_stamp = int(time.mktime(date_time.timetuple()))
    return time_stamp


def stamp2date(timestamp):
    """
    :param timestamp: int timestamp
    :return: datetime
    """
    date_time = datetime.datetime.utcfromtimestamp(timestamp)
    return date_time


def sum_dict(keys: Union[list, tuple, set], data: Union[list, tuple, set]):

    result = dict()
    for element in data:
        for key in keys:
            res_key = "total_" + key
            result[res_key] = result.get(res_key, 0.00) + element.get(key)
    return result


