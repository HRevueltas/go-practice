function loops() {
    let i = 1;
    // for loop for a single condition
    while (i < 3) {
        console.log(i);
        i++;
    }

    // for loop for a classic initial/condition/after
    for (let j = 0; j <= 3; j++) {
        console.log(j);
    }

    // for loop with break
    while (true) {
        console.log("loop");
        break;
    }

    // for loop with continue
    for (let n = 0; n < 6; n++) {
        if (n % 2 === 0) {
            continue;
        }
        console.log(n);
    }
}

function main() {
    // record the start time
    const start = new Date();

    // call the loops function
    loops();

    // record the end time
    const end = new Date();

    // calculate the duration
    const duration = end - start;

    // print the duration
    console.log("Duration: ", duration, "ms");
}

// Call the main function
main();
