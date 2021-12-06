package schedule_max_jobs

func ScheduleMaxJobs(intervals [][2]int) [][2]int {
	intervals = SortIntervalsByEnd(intervals)
	lastInterval := intervals[0]
	maxIntervals := [][2]int { lastInterval }

	for i := 1; i < len(intervals); i++ {
		if (lastInterval[1] >= intervals[i][0]) { continue }
		maxIntervals = append(maxIntervals, intervals[i])
		lastInterval = intervals[i]
	}

  return maxIntervals
}

func SortIntervalsByEnd(intervals [][2]int) [][2]int  {
	isSorted := true

	for i := 0; i < len(intervals) - 1; i++ {
		a := intervals[i]
		b := intervals[i + 1]

		if b[1] < a[1] {
			intervals[i] = b
			intervals[i + 1] = a

			isSorted = false
		}
	}

	if isSorted == false {
		return SortIntervalsByEnd(intervals)
	} else {
		return intervals
	}
}

/*



Problem: Movie Scheduling Problem
Input: A set I of n intervals on the line.
[][2]int{{4, 12}, {2, 9}, {10, 15}, {6, 15}, {14, 34}, {16, 20}, {21, 30}, {22, 30}, {28, 46}, {32, 48}}

Output: What is the largest subset of mutually non-overlapping intervals that can be selected from I?
[][2]int{{2, 9}, {10, 15}, {16, 20}, {22, 30}, {32, 48}}

	isSorted = true
	iterate over array from 0 to second to last (< len - 1)
		compare i to i + 1
		if i+1 < i, switch position, isSorted = false

	if isSort == false, call func and return that
	else return sorted

	- save lastInterval
	- save largestSetNonOverlapIntervals
	- iterate over intervals from 1 to len
		- if (lastInterval end >= i start) { skip }
		- append interval
		- set last interval
	- return largest

Input: [[0, 10], [5, 10]] // set of projects with start and end times
Output: [[]...] // largest possible set of non-overlapping projects.

Approach: Full Verification through recursive branch check

Algo:
- Create array of possible job sets
- Create array for temp job set

call create possible job sets(all jobs array)
  - If no more jobs
      - add job set to possible job sets
      - reset job set to empty
  - Pick the project that is most recent
      - add to temp job set
      - call create possible job sets without it

- array with all possible job sets
  - map/reduce to number of jobs
  - select first job set with maximum number


	heuristic
	- sort by interval end
		1 2 8 4 6 0
		0 1 2 3 4 5





*/


/*

Input:
  slice of arrays
  each subarray represents start and end time for job
  
Output:
  largest subset of mutually non-overlapping intervals
  the most jobs that don't overlap
  
Examples:
  in general, start with earliest start time
  use shortest jobs
  can't overlap
  
{2,8}, {2,10}, {3,5}, {3,8} {3,9}
  
Data:
  need to compare subarrays
  append to result slice
  
  
General Approaches:
#1: for each job, calculate value of choosing that job, pick one. Then for remaining jobs, give weight to each, select one. Repeat until none left.

Weight => the job with soonest end date that doesn't overlap
shortest distance to end heuristic?

  
Algorithm:
  look for earliest start time
    of jobs with that start time, choose the shortest
    see if any other job falls within that range
      and end time is before end time of current shortest job
    
Algorithm (2):
- (Sort by start time and then duration)

- INPUT: slice of jobs
- OUTPUT: the longest slice of non-conflicting jobs selected from the input

- For each job in the input slice:
  - nonConflicting = Create a new subarray of jobs which do not conflict
  - if the length of nonConflicting is 0:
    - return a slice with the job in it
  - if the length of nonConflicting is 1:
    - return a slice with the job and the one nonConflicting job
  - else:
    - for each job in nonConflicting, call this function again and put the resulting return in an slice (jobSets)
    - return a slice with the current job appended to longest subarray in jobSets
    
    
{1, 5}, {2, 3}, {4, 10}, {6, 7}, {8, 9}
-> {1, 5}: nonConflicting: {6, 7}, {8, 9}
   -> {6, 7}: nonC: {8, 9}
      RETURN: [{6, 7}, {8, 9}]
   -> {8, 9}: nonC: {6, 7}
   RETURN: [[1, 5], [6,7], [8,9]]
   
-> {2, 3}:  nonC: {4, 10}, {6, 7}, {8, 9}
   -> {4, 10}: nonC: []
   -> {6, 7}: [{8,9}]
      RETURN [{6, 7}, {8, 9}]
   RETURN: {2, 3}, {6,7}, {8, 9}

-> {4, 10}...

Now consider the following scheduling problem. Imagine you are a highly in
demand actor, who has been presented with offers to star in n different movie
projects under development. Each offer comes specified with the first and last
day of filming. Whenever you accept a job, you must commit to being available
throughout this entire period. Thus, you cannot accept two jobs whose intervals
overlap.

For an artist such as yourself, the criterion for job acceptance is clear: you
want to make as much money as possible. Because each film pays the same fee,
this implies you seek the largest possible set of jobs (intervals) such that no two
of them conflict with each other.

For example, consider the available projects in Figure 1.5. You can star in
at most four films, namely “Discrete” Mathematics, Programming Challenges,
Calculated Bets, and one of either Halting State or Steiner’s Tree.

You (or your agent) must solve the following algorithmic scheduling problem:

Problem: Movie Scheduling Problem
Input: A set I of n intervals on the line.
[][2]int{{4, 12}, {2, 9}, {10, 15}, {6, 15}, {14, 34}, {16, 20}, {21, 30}, {22, 30}, {28, 46}, {32, 48}}

Output: What is the largest subset of mutually non-overlapping intervals that
can be selected from I?
[][2]int{{2, 9}, {10, 15}, {16, 20}, {22, 30}, {32, 48}}

Now you (the algorist) are given the job of developing a scheduling algorithm
for this task. Stop right now and try to find one. Again, I’ll be happy to wait. . .

end = start => conflicting

*/
