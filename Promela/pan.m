#define rand	pan_rand
#define pthread_equal(a,b)	((a)==(b))
#if defined(HAS_CODE) && defined(VERBOSE)
	#ifdef BFS_PAR
		bfs_printf("Pr: %d Tr: %d\n", II, t->forw);
	#else
		cpu_printf("Pr: %d Tr: %d\n", II, t->forw);
	#endif
#endif
	switch (t->forw) {
	default: Uerror("bad forward move");
	case 0:	/* if without executable clauses */
		continue;
	case 1: /* generic 'goto' or 'skip' */
		IfNotBlocked
		_m = 3; goto P999;
	case 2: /* generic 'else' */
		IfNotBlocked
		if (trpt->o_pm&1) continue;
		_m = 3; goto P999;

		 /* PROC LinearRegression */
	case 3: // STATE 1 - parcial.pml:14 - [i = 0] (0:0:1 - 1)
		IfNotBlocked
		reached[0][1] = 1;
		(trpt+1)->bup.oval = now.i;
		now.i = 0;
#ifdef VAR_RANGES
		logval("i", now.i);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 4: // STATE 2 - parcial.pml:14 - [((i<=(10-1)))] (0:0:0 - 1)
		IfNotBlocked
		reached[0][2] = 1;
		if (!((now.i<=(10-1))))
			continue;
		_m = 3; goto P999; /* 0 */
	case 5: // STATE 3 - parcial.pml:15 - [((mutex>0))] (6:0:1 - 1)
		IfNotBlocked
		reached[0][3] = 1;
		if (!((now.mutex>0)))
			continue;
		/* merge: mutex = (mutex-1)(0, 4, 6) */
		reached[0][4] = 1;
		(trpt+1)->bup.oval = now.mutex;
		now.mutex = (now.mutex-1);
#ifdef VAR_RANGES
		logval("mutex", now.mutex);
#endif
		;
		_m = 3; goto P999; /* 1 */
	case 6: // STATE 6 - parcial.pml:16 - [sumX = (sumX+X[i])] (0:0:1 - 1)
		IfNotBlocked
		reached[0][6] = 1;
		(trpt+1)->bup.oval = now.sumX;
		now.sumX = (now.sumX+now.X[ Index(now.i, 10) ]);
#ifdef VAR_RANGES
		logval("sumX", now.sumX);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 7: // STATE 7 - parcial.pml:17 - [sumY = (sumY+Y[i])] (0:0:1 - 1)
		IfNotBlocked
		reached[0][7] = 1;
		(trpt+1)->bup.oval = now.sumY;
		now.sumY = (now.sumY+now.Y[ Index(now.i, 10) ]);
#ifdef VAR_RANGES
		logval("sumY", now.sumY);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 8: // STATE 8 - parcial.pml:18 - [sumXY = (sumXY+(X[i]*Y[i]))] (0:0:1 - 1)
		IfNotBlocked
		reached[0][8] = 1;
		(trpt+1)->bup.oval = now.sumXY;
		now.sumXY = (now.sumXY+(now.X[ Index(now.i, 10) ]*now.Y[ Index(now.i, 10) ]));
#ifdef VAR_RANGES
		logval("sumXY", now.sumXY);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 9: // STATE 9 - parcial.pml:19 - [sumXX = (sumXX+(X[i]*X[i]))] (0:0:1 - 1)
		IfNotBlocked
		reached[0][9] = 1;
		(trpt+1)->bup.oval = now.sumXX;
		now.sumXX = (now.sumXX+(now.X[ Index(now.i, 10) ]*now.X[ Index(now.i, 10) ]));
#ifdef VAR_RANGES
		logval("sumXX", now.sumXX);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 10: // STATE 10 - parcial.pml:20 - [mutex = (mutex+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][10] = 1;
		(trpt+1)->bup.oval = now.mutex;
		now.mutex = (now.mutex+1);
#ifdef VAR_RANGES
		logval("mutex", now.mutex);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 11: // STATE 11 - parcial.pml:14 - [i = (i+1)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][11] = 1;
		(trpt+1)->bup.oval = now.i;
		now.i = (now.i+1);
#ifdef VAR_RANGES
		logval("i", now.i);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 12: // STATE 17 - parcial.pml:24 - [m = (((10*sumXY)-(sumX*sumY))/((10*sumXX)-(sumX*sumX)))] (0:0:1 - 3)
		IfNotBlocked
		reached[0][17] = 1;
		(trpt+1)->bup.oval = ((P0 *)_this)->m;
		((P0 *)_this)->m = (((10*now.sumXY)-(now.sumX*now.sumY))/((10*now.sumXX)-(now.sumX*now.sumX)));
#ifdef VAR_RANGES
		logval("LinearRegression:m", ((P0 *)_this)->m);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 13: // STATE 18 - parcial.pml:26 - [b = ((sumY-(m*sumX))/10)] (0:0:1 - 1)
		IfNotBlocked
		reached[0][18] = 1;
		(trpt+1)->bup.oval = ((P0 *)_this)->b;
		((P0 *)_this)->b = ((now.sumY-(((P0 *)_this)->m*now.sumX))/10);
#ifdef VAR_RANGES
		logval("LinearRegression:b", ((P0 *)_this)->b);
#endif
		;
		_m = 3; goto P999; /* 0 */
	case 14: // STATE 19 - parcial.pml:26 - [printf('Coeficiente m: %d, Coeficiente b: %d\\n',m,b)] (0:0:0 - 1)
		IfNotBlocked
		reached[0][19] = 1;
		Printf("Coeficiente m: %d, Coeficiente b: %d\n", ((P0 *)_this)->m, ((P0 *)_this)->b);
		_m = 3; goto P999; /* 0 */
	case 15: // STATE 20 - parcial.pml:27 - [-end-] (0:0:0 - 1)
		IfNotBlocked
		reached[0][20] = 1;
		if (!delproc(1, II)) continue;
		_m = 3; goto P999; /* 0 */
	case  _T5:	/* np_ */
		if (!((!(trpt->o_pm&4) && !(trpt->tau&128))))
			continue;
		/* else fall through */
	case  _T2:	/* true */
		_m = 3; goto P999;
#undef rand
	}

