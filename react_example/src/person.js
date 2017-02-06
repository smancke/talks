
class Person {
    constructor(first, last) {
        this.first  = first
        this.last = last
    }

    get name() {
        return `${this.first} ${this.last}`
    }
}

export default Person
