export default class ArrayList<T> {
    public length: number;
    private array: T[];
    private capacity: number;

    constructor(capacity: number) {
        this.capacity = capacity;
        this.length = 0;
        this.array = new Array(this.capacity);
    }

    prepend(item: T): void {
        this.length++;

        // Increase the capacity of the array if current capacity is reached.
        if (this.length > this.capacity) {
            this.capacity = this.capacity * 2;
            let newArray = new Array(this.capacity);
            newArray[0] = item;

            // Copy the data over to the new array. And prepend the new item.
            for (let i = 1; i < this.length; i++) {
                newArray[i] = this.array[i - 1];
            }

            this.array = newArray;

            return;
        }

        // Shift the array one space to the right.
        for (let i = this.length - 1; i >= 0; i--) {
            this.array[i + 1] = this.array[i];
        }

        this.array[0] = item;
    }

    insertAt(item: T, idx: number): void {}

    append(item: T): void {
        this.length++;

        // Increase the capacity of the array if current capacity is reached.
        if (this.length > this.capacity) {
            this.capacity = this.capacity * 2;
            let newArray = new Array(this.capacity);

            // Copy the data over to the new array.
            for (let i = 0; i < this.length; i++) {
                newArray[i] = this.array[i];
            }

            this.array = newArray;
        }

        this.array[this.length - 1] = item;
    }

    remove(item: T): T | undefined {
        let removedIdx: number | undefined = undefined;

        for (let i = 0; i < this.length; i++) {
            if (this.array[i] === item) {
                removedIdx = i;
                break;
            }
        }				

        if (removedIdx !== undefined) {
            return this.removeAt(removedIdx);
        }

        return;
    }

    get(idx: number): T | undefined {
        return this.array[idx];
    }

    removeAt(idx: number): T | undefined {
        if (idx >= this.length) {
            return undefined;
        }

        const value = this.array[idx];

        // Shift the values in the array from the index to the left.
        for (let i = idx; i < this.length; i++) {
            this.array[i] = this.array[i + 1];
        }

        this.length--;

        return value;
    }
}
