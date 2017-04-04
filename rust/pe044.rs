// https://projecteuler.net/problem=44

use std::collections::HashMap;

fn pengaton(n: i32) -> i32 {
    n * (3 * n - 1) / 2
}

fn main() {
    // Create pentagon numbers
    let maxval = 10000;
    let mut pents_map = HashMap::new();
    let mut pents_rev = HashMap::new();
    for n in 1..maxval {
        let buff = pengaton(n);
        pents_map.insert(n, buff);
        pents_rev.insert(buff, n);
    }

    let mut min_d: i32;
    match pents_map.get(&(maxval-1)) {
        Some(&number) => min_d = number,
        _ => min_d=-1,
    }

    for i in 1..maxval {
        for k in 1..maxval {
            let j = i + k;
            if j >= maxval {
                continue
            }
            let pi;
            let pj;
            match pents_map.get(&i) {
                Some(&number) => pi=number,
                _ => continue,
            }
            match pents_map.get(&j) {
                Some(&number) => pj=number,
                _ => continue,
            }
            let sum = pi+pj;
            let del = pj-pi;
            if pents_rev.contains_key(&sum) && pents_rev.contains_key(&del) {
                if del < min_d {
                    min_d = del;
                }
            }
        }
        println!("i={}", i);
    }
    println!("The answer is {}", min_d);
}
