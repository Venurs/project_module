# Generated by Django 2.1.7 on 2019-04-30 06:47

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('account', '0001_initial'),
    ]

    operations = [
        migrations.AlterModelOptions(
            name='kjzduser',
            options={'managed': False, 'permissions': (('can_login_cms', 'Can Login Cms'),)},
        ),
    ]
