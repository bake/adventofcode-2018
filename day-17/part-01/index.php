<?php
// Why did i write this in PHP?

error_reporting(E_ALL);
ini_set('display_errors', 1);
ini_set('memory_limit', '-1');

const START = '+';
const SAND = '.';
const CLAY = '#';
const FLOW = '|';
const STILL = '~';

class Point
{
	public $x, $y;

	public function __construct(int $x, int $y)
	{
		$this->x = $x;
		$this->y = $y;
	}

	public function __toString() : string
	{
		return "($this->x,$this->y)";
	}
}

class Map
{
	private $map;
	private $iteration = 0;

	public function __construct($map)
	{
		$this->map = $map;
	}

	public function at(int $x, int $y) : string
	{
		if (!isset($this->map[$y][$x])) {
			return SAND;
		}
		if ($y < 0 || $y >= count($this->map)) {
			return CLAY;
		}
		if ($x < 0 || $x >= count($this->map[0])) {
			return CLAY;
		}
		return $this->map[$y][$x];
	}

	public function is(int $x, int $y, string ...$types) : bool
	{
		return in_array($this->at($x, $y), $types);
	}

	public function isSolid(int $x, int $y) : bool
	{
		return $this->is($x, $y, CLAY, STILL);
	}

	public function set(int $x, int $y, string $value) : bool
	{
		if ($y < 0 || $y >= count($this->map)) {
			return false;
		}
		if ($x < 0 || $x >= count($this->map[0])) {
			return false;
		}
		$this->map[$y][$x] = $value;
		return true;
	}

	public function print()
	{
		foreach ($this->map as $y => $row) {
			foreach ($row as $x => $cell) {
				echo $cell;
			}
			echo PHP_EOL;
		}
	}

	public function image()
	{
		$minX = INF;
		for ($y = 0; $y < count($this->map); $y++) {
			for ($x = 0; $x < count($this->map[$y]); $x++) {
				if ($this->is($x, $y, CLAY) && $x < $minX) {
					$minX = $x;
				}
			}
		}

		$image = imagecreate(count($this->map[0]) - $minX, count($this->map));
		if ($image === false) {
			exit;
		}
		$colors = [
			START => imagecolorallocate($image, 255, 0, 0),
			CLAY => imagecolorallocate($image, 0, 0, 0),
			SAND => imagecolorallocate($image, 255, 255, 255),
			FLOW => imagecolorallocate($image, 119, 181, 254),
			STILL => imagecolorallocate($image, 0, 178, 228),
		];
		for ($y = 0; $y < count($this->map); $y++) {
			for ($x = 0; $x < count($this->map[$y]); $x++) {
				imagesetpixel($image, $x, $y, $colors[$this->at($x + $minX, $y)]);
			}
		}
		return $image;
	}

	public function simulate(int $x, int $y)
	{
		while ($this->flow($x, $y));
	}

	private function flow(int $x, int $y) : bool
	{
		// echo '.';
		// echo "Iteration $this->iteration\n";
		// imagepng($this->image(), "out/$this->iteration.png");
		// $this->iteration++;

		for ($dy = 0; $y + $dy < count($this->map) && !$this->isSolid($x, $y + $dy); $dy++) {
			$this->set($x, $y + $dy, FLOW);
		}
		$y += $dy - 1;

		$settled = false;
		$flew = false;
		for ($dl = 0; $dl < count($this->map); $dl--) {
			// echo ',';
			if ($y > count($this->map)) {
				return false;
			}
			if ($this->isSolid($x + $dl, $y)) {
				break;
			}
			$this->set($x + $dl, $y, FLOW);
			if ($this->isSolid($x + $dl, $y + 1)) {
				continue;
			}
			$settled |= $this->flow($x + $dl, $y);
			$flew = true;
			break;
		}

		for ($dr = 0; $dr < count($this->map); $dr++) {
			if ($this->isSolid($x + $dr, $y)) {
				break;
			}
			$this->set($x + $dr, $y, FLOW);
			if ($this->isSolid($x + $dr, $y + 1)) {
				continue;
			}
			$settled |= $this->flow($x + $dr, $y);
			$flew = true;
			break;
		}

		if (!$flew) {
			for ($dx = $dl + 1; $dx < $dr; $dx++) {
				$this->set($x + $dx, $y, STILL);
			}
		}
		return $settled;
	}

	private function flowDown(int $x, int $y)
	{
		switch ($this->at($x, $y + 1)) {
			case SAND:
			case FLOW:
				$this->set($x, $y + 1, FLOW);
				$this->flowDown($x, $y + 1);
				break;
			case CLAY:
			case STILL:
				return;
		}
	}
}

$file = new SplFileObject('../input.txt');
$grid = [];
[$minX, $maxX] = [INF, -INF];
[$minY, $maxY] = [INF, -INF];
while (!$file->eof()) {
	preg_match('/^(x|y)=([0-9]+), (x|y)=([0-9]+)\.\.([0-9]+)$/', $file->fgets(), $matches);
	$fields = [
		$matches[1] => [$matches[2]],
		$matches[3] => range($matches[4], $matches[5]),
	];
	foreach ($fields['y'] as $y) {
		foreach ($fields['x'] as $x) {
			if ($x > $maxX) {
				$maxX = $x;
			}
			if ($x < $minX) {
				$minX = $x;
			}
			if ($y > $maxY) {
				$maxY = $y;
			}
			if ($y < $minY) {
				$minY = $y;
			}
			$grid[$y][$x] = CLAY;
		}
	}
}
for ($y = 0; $y < $maxY; $y++) {
	if (!isset($grid[$y])) {
		$grid[$y] = array_fill(0, $maxX, SAND);
		continue;
	}
	for ($x = 0; $x < $maxX; $x++) {
		if (!isset($grid[$y][$x])) {
			$grid[$y][$x] = SAND;
		}
	}
}

$map = new Map($grid);
for ($i = 0; $i < 82; $i++) {
	echo "$i ";
	$map->simulate(500, 0);
	echo PHP_EOL;
}
imagepng($map->image(), 'out.png');
