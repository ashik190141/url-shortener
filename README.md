**Tech Stack**

* Go
* net/http
* gorilla/mux
* GORM
* PostgreSQL

### **Snowflake: 64-bit Distributed ID**

```text
| 1 bit  | 41 bits      | 5 bits       | 5 bits     | 12 bits  |
| sign=0 | timestamp ms | datacenter   | machine ID | sequence |
```

* **41-bit timestamp** → ~69 years of IDs
* **10-bit node identity** → 1,024 unique nodes
* **12-bit sequence** → 4,096 IDs per millisecond per node
* **Result:** 4M+ unique IDs/sec across the cluster with zero coordination.

### system desugn of create a url shortener

![System Design](https://ibb.co.com/XZLNJC8g)
