package geom

/*
Conventions:
Right-Handed coordinate system
	Rotation around thumb axis, fingers curled
Column-Major
	| Xx Yx Zx Tx |   | X |
	| Xy Yy Zy Ty | * | Y |
	| Xz Yz Zz Tz |   | Z |
	| 0  0  0  1  |   | 1 |

       Z
      /
     /
    ----- X
    |
    |
    |
    Y

    Roll Pitch Yaw -> Ry, Rx, Rz

    a.Product(b) -> a is performed first
*/
