# coding=utf8
# Create your views here.

from django.conf.urls import include
from django.urls import path
from rest_framework import routers
from account import apis


urlpatterns = [
    path('log_in/', apis.log_in),
    path('log_out/', apis.log_out),
]



