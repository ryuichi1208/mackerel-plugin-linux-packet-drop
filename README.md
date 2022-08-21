# mackerel-plugin-linux-packet-drop

A mackerel plugin for measuring Linux packet loss

## Graphs and Metrics

* ${devicename}.RxDropped
* ${devicename}.RxErrors
* ${devicename}.TxDropped
* ${devicename}.TxErrors

```
PacketDrop.docker0.docker0.RxDropped	0	1661064368
PacketDrop.docker0.docker0.TxDropped	0	1661064368
PacketDrop.docker0.docker0.RxErrors	0	1661064368
PacketDrop.docker0.docker0.TxErrors	0	1661064368
PacketDrop.vethcec3d996.vethcec3d996.RxDropped	0	1661064368
PacketDrop.vethcec3d996.vethcec3d996.TxDropped	0	1661064368
PacketDrop.vethcec3d996.vethcec3d996.RxErrors	0	1661064368
PacketDrop.vethcec3d996.vethcec3d996.TxErrors	0	1661064368
PacketDrop.lo.lo.RxDropped	0	1661064368
PacketDrop.lo.lo.TxDropped	0	1661064368
PacketDrop.lo.lo.RxErrors	0	1661064368
PacketDrop.lo.lo.TxErrors	0	1661064368
PacketDrop.ens160.ens160.RxDropped	60	1661064368
PacketDrop.ens160.ens160.TxDropped	0	1661064368
PacketDrop.ens160.ens160.RxErrors	0	1661064368
PacketDrop.ens160.ens160.TxErrors	0	1661064368
PacketDrop.nodelocaldns.nodelocaldns.RxDropped	0	1661064368
PacketDrop.nodelocaldns.nodelocaldns.TxDropped	0	1661064368
PacketDrop.nodelocaldns.nodelocaldns.RxErrors	0	1661064368
PacketDrop.nodelocaldns.nodelocaldns.TxErrors	0	1661064368
PacketDrop.flannel.1.flannel.1.RxDropped	0	1661064368
PacketDrop.flannel.1.flannel.1.TxDropped	0	1661064368
PacketDrop.flannel.1.flannel.1.RxErrors	0	1661064368
PacketDrop.flannel.1.flannel.1.TxErrors	0	1661064368
PacketDrop.cni0.cni0.RxDropped	0	1661064368
PacketDrop.cni0.cni0.TxDropped	0	1661064368
PacketDrop.cni0.cni0.RxErrors	0	1661064368
PacketDrop.cni0.cni0.TxErrors	0	1661064368
PacketDrop.vethe0705fed.vethe0705fed.RxDropped	0	1661064368
PacketDrop.vethe0705fed.vethe0705fed.TxDropped	0	1661064368
PacketDrop.vethe0705fed.vethe0705fed.RxErrors	0	1661064368
PacketDrop.vethe0705fed.vethe0705fed.TxErrors	0	1661064368
```

## Reference

[net/core/net-procfs.c](https://elixir.bootlin.com/linux/latest/source/net/core/net-procfs.c)
