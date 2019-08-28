# coding=utf-8

from django.db import models
from django.contrib.auth.models import BaseUserManager, AbstractBaseUser, PermissionsMixin

from common.models import DeleteStatusMixin, GmtCreateModifiedTimeMixin


class MyUserManager(BaseUserManager):
    def create_user(self, phone, name, password=None):
        """
        Creates and saves a User with the given phone, date of
        birth and password.
        """

        if not phone:
            raise ValueError('Users must have a phone number')

        user = self.model(
            # email=self.normalize_email(email),
            name=name,
            phone=phone,
        )
        user.set_password(password)
        user.save(using=self._db)
        return user

    def create_superuser(self, phone, name, password):
        """
        Creates and saves a superuser with the given phone, date of
        birth and password.
        """
        user = self.create_user(
            phone=phone,
            # email=email,
            password=password,
            name=name,
        )
        user.is_admin = True
        user.is_superuser = True
        user.save(using=self._db)
        return user


class MyUser(AbstractBaseUser, PermissionsMixin):
    email = models.EmailField(
        verbose_name='email address',
        max_length=255,
    )
    name = models.CharField(
        verbose_name='name',
        max_length=150,
        blank=True,
        null=True,
        error_messages={
            'unique': "A user with that email already exists.",
        },
    )

    phone = models.CharField(
        verbose_name='telephone number',
        max_length=64,
        blank=True,
        unique=True,
        null=True,
        error_messages={
            'unique': "A user with that telephone already exists.",
        },
    )
    is_active = models.BooleanField(default=True)
    is_admin = models.BooleanField(default=False)

    kjzd_user_id = models.IntegerField(unique=True, null=True)

    objects = MyUserManager()

    USERNAME_FIELD = 'phone'
    REQUIRED_FIELDS = ['name']

    def __str__(self):
        return self.phone

    @property
    def is_staff(self):
        """ Is the user a member of staff? """
        # Simplest possible answer: All admins are staff
        return self.is_admin


class KjzdUser(models.Model):

    password = models.CharField(max_length=128)
    nick = models.CharField(unique=True, max_length=30, blank=True, null=True)
    email = models.CharField(unique=True, max_length=254, blank=True, null=True)
    tel = models.CharField(unique=True, max_length=64, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'pre_users'
        permissions = (
            ("can_login_cms", "Can Login Cms"),
        )


class Salesman(GmtCreateModifiedTimeMixin, DeleteStatusMixin):

    name = models.CharField(max_length=128, blank=True, null=True, verbose_name='销售姓名')

    objects = models.Manager()