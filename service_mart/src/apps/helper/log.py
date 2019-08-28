import logging


def get_logger(name, prefix='service_mart'):
    return logging.getLogger('{prefix}.{name}'.format(prefix=prefix, name=name))


class ExcludeFilter(logging.Filter):
    def __init__(self, exclude_name='service_mart.actions', exclude_level=logging.ERROR):
        super(ExcludeFilter, self).__init__()
        self.exclude_name = exclude_name
        self.exclude_level = exclude_level

    def filter(self, record):
        return not (record.name == self.exclude_name or record.levelno >= self.exclude_level)

