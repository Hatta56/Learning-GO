<?php

function getDigit($N, $d)
{
	$string = strval($N);
	return $string[$d - 1];
}

function getNthChar($N)
{

	$sum = 0;
	$ketik = 1111;

	$dist = 0;

	for($len = 1; $len < $N; $len++)
	{

		$sum += $ketik * $len;
		$dist += $ketik;
		if ($sum >= $N)
		{
			$sum -= $ketik * $len;
			$dist -= $ketik;
			$N -= $sum;
			break;
		}
		$ketik *= 10;
	}

	$diff = ($N / $len) + 1;

	$d = $N % $len;
	if ($d == 0)
		$d = $len;

	return getDigit($dist + $diff, $d);
}


$N = 1000;
echo getNthChar($N);


?>
