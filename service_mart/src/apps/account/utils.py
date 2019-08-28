# coding=utf-8
import hashlib


def str2md5(original_str):

    """ 字符串md5加密,用于登陆时比对密码
    :param original_str: 原始字符串
    :return:
    """

    md_str = hashlib.md5()
    md_str.update(original_str.encode(encoding='utf-8'))
    md_str = md_str.hexdigest()
    md_str = md_str.upper()
    md_str = md_str[::-1]
    md_str = md_str.lower()

    new_md_str = hashlib.md5()
    new_md_str.update(md_str.encode(encoding='utf-8'))
    new_md_str = new_md_str.hexdigest()

    return new_md_str[:32]
