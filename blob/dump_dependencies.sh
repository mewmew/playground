#Dynamic section at offset 0x4bf0a0 contains 19 entries:
#  Tag        Type                         Name/Value
# 0x0000000000000001 (NEEDED)             Shared library: [libpthread.so.0]
# 0x0000000000000001 (NEEDED)             Shared library: [libc.so.6]
# 0x0000000000000004 (HASH)               0x8be1a8
# 0x0000000000000006 (SYMTAB)             0x8be768
# 0x000000000000000b (SYMENT)             24 (bytes)
# 0x0000000000000005 (STRTAB)             0x8be330
# 0x000000000000000a (STRSZ)              320 (bytes)
# 0x0000000000000007 (RELA)               0x8be120
# 0x0000000000000008 (RELASZ)             24 (bytes)
# 0x0000000000000009 (RELAENT)            24 (bytes)
# 0x0000000000000003 (PLTGOT)             0x8bf008
# 0x0000000000000014 (PLTREL)             RELA
# 0x0000000000000002 (PLTRELSZ)           384 (bytes)
# 0x0000000000000017 (JMPREL)             0x8be5e8
# 0x0000000000000015 (DEBUG)              0x0
# 0x000000006ffffffe (VERNEED)            0x8be168
# 0x000000006fffffff (VERNEEDNUM)         2
# 0x000000006ffffff0 (VERSYM)             0x8be138
# 0x0000000000000000 (NULL)               0x0

printf "=== [ shared library ] ===\n"
printf "\n"
readelf -d $1 | grep "Shared library:" | cut -b 58- | sort | uniq
printf "\n"

#Symbol table '.dynsym' contains 21 entries:
#   Num:    Value          Size Type    Bind   Vis      Ndx Name
#     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND
#     1: 000000000057804f     0 OBJECT  GLOBAL DEFAULT   11 crosscall2
#     2: 0000000000577d39     0 FUNC    GLOBAL DEFAULT   11 _cgo_allocate
#     3: 0000000000577de7     0 FUNC    GLOBAL DEFAULT   11 _cgo_panic
#     4: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND getaddrinfo@GLIBC_2.2.5 (3)
#     5: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND free@GLIBC_2.2.5 (3)
#     6: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND gai_strerror@GLIBC_2.2.5 (3)
#     7: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND __errno_location@GLIBC_2.2.5 (2)
#     8: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND freeaddrinfo@GLIBC_2.2.5 (3)
#     9: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND pthread_attr_init@GLIBC_2.2.5 (2)
#    10: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND pthread_attr_getstacksize@GLIBC_2.2.5 (2)
#    11: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND pthread_attr_destroy@GLIBC_2.2.5 (2)
#    12: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND sigfillset@GLIBC_2.2.5 (3)
#    13: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND sigprocmask@GLIBC_2.2.5 (3)
#    14: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND pthread_create@GLIBC_2.2.5 (2)
#    15: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND strerror@GLIBC_2.2.5 (3)
#    16: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND stderr@GLIBC_2.2.5 (3)
#    17: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND fprintf@GLIBC_2.2.5 (3)
#    18: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND abort@GLIBC_2.2.5 (3)
#    19: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND malloc@GLIBC_2.2.5 (3)
#    20: 0000000000000000     0 OBJECT  GLOBAL DEFAULT  UND fwrite@GLIBC_2.2.5 (3)

printf "=== [ functions ] ===\n"
printf "\n"
readelf --dyn-syms $1 | grep " UND " | grep -v " LOCAL " | cut -b 60- | sort | uniq
