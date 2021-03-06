"""service_mart URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

admin.autodiscover()

urlpatterns = [
    path('smapi/admin/', admin.site.urls),
    path('smapi/account/api/', include('account.urls'), name='account'),
    path('smapi/ordering/api/', include('ordering.urls'), name='ordering'),
    path('smapi/appeal/api/', include('appeal.urls'), name='appeal'),
    path('smapi/promote/api/', include('promote.urls'), name='promote'),
    path('smapi/wechatpay/api/', include('wechatpay.urls'), name='wechatpay'),
    path('smapi/upload/api/', include('upload.urls'), name='upload')
]
