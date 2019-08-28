# coding=utf-8
from django.db import models
from django.db.models import Q
from django.utils import timezone
from django.utils.translation import ugettext_lazy as _


class GmtCreateModifiedTimeMixin(models.Model):
    gmt_create = models.DateTimeField(_('gmt create'), auto_now_add=True)

    gmt_modified = models.DateTimeField(_('gmt modified'), auto_now=True)

    class Meta:
        abstract = True


class DeleteStatusMixinQuerySet(models.QuerySet):
    def filter_deleted(self):
        return self.filter(delete_status=DeleteStatusMixin.STATUS_DELETED)

    def filter_not_deleted(self):
        return self.filter(delete_status=DeleteStatusMixin.STATUS_NOT_DELETED)


class DeleteStatusMixin(models.Model):
    STATUS_NOT_DELETED = 0
    STATUS_DELETED = 1

    DELETE_STATUS_CHOICE = (
        (STATUS_NOT_DELETED, '未删除'),
        (STATUS_DELETED, '已删除')
    )

    delete_status = models.PositiveSmallIntegerField(
        _('删除状态'), choices=DELETE_STATUS_CHOICE, blank=True, null=False, default=STATUS_NOT_DELETED
    )

    delete_status_manager = DeleteStatusMixinQuerySet.as_manager()

    class Meta:
        abstract = True


class StartTimeEndTimeMixinQuerySet(models.QuerySet):
    def filter_between_start_time_and_end_time(self, datetime_to_compare=None):
        if not datetime_to_compare:
            datetime_to_compare = timezone.now()
        return self.filter(Q(start_time=None) | Q(start_time__lte=datetime_to_compare),
                           Q(end_time=None) | Q(end_time__gt=datetime_to_compare))


class StartTimeEndTimeMixin(models.Model):
    start_time = models.DateTimeField(_('start time'), blank=True, null=True)
    end_time = models.DateTimeField(_('end time'), blank=True, null=True)

    start_time_end_time_manager = StartTimeEndTimeMixinQuerySet.as_manager()

    class Meta:
        abstract = True

    @property
    def is_between_start_time_and_end_time(self):
        now = timezone.now()
        return (not self.start_time or self.start_time <= now) and (not self.end_time or self.end_time > now)