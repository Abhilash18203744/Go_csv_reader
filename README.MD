# customerimporter
Package is provided with few unit tests (in interview_test.go file) and required data files (customer.csv, customer_wrong.csv).

## Code Structure
  Functions written in the package takes csv file path as an input and returns a sorted map of domain name as key and respective users count as value.
  
  Two separate functions are included in the package :- 1) basicDomainCounter() 2) concWPDomainCounter()

  1) basicDomainCounter() function processes one row at a time (which includes fetching email, validating it with the help of'@', return domain name). As the line in csv file is short (i.e. only 5 columns) and fetching domain name is simple operation, this function performs well for small size as well as large size input csv file. But for complicated processes (e.g. complex parsing operations or expensive tasks) performance will degrade. 
  
  2) concWPDomainCounter() function uses concurrency and worker pool technique to improve the performance in case any complicated parsing operations or expensive tasks are included in the functionality. In this function worker pool is implemented using goroutines and channels. Worker function is implemented to run several concurrent instances. Workers will receive work on the jobs channel (with buffer 200 and changed as per requirement) and send corresponding output on results channel. Wait group is used to wait for created multiple goroutines to finish.

  Both the functions are added in the package for the comparison of the performance, and several validations are done for input data e.g. if the input email address is wrong code will skip that entry without exit. For testing purpose a sleep statement was added in the parsing function in order to produce expensive task, and test concWPDomainCounter() function, which worked well as expected. Functions are written assuming that only column named 'email' contain email address, subsequently domain names. Also scenario of missing 'email' column is handled.
 
## Test cases
  Following scenarios are considered for unit testing which includes positive as well as negative test cases:
  1) Passing an empty input csv file.
  2) Passing a csv file with no email records in it.
  3) Passing wrong file location/ file name.
  4) Passing a correct small csv file.
  5) Passing a correct large csv file.
  