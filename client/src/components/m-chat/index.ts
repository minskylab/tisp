import { html, define, Hybrids } from "hybrids";
import styles from "./index.css";

// const styles = `
//   div.wrapper {
//     background-color: aquamarine;
//   }
// `;

console.log(styles);
interface SimpleCounter extends HTMLElement {
  count: number;
}

export function increaseCount(host: SimpleCounter) {
  host.count += 1;
}

export const SimpleCounter: Hybrids<SimpleCounter> = {
  count: 0,
  render: ({ count }) =>
    html`
      <div class=${styles.wrapper}>
        <div>hello world: ${count}</div>
        <div>
          <button onclick="${increaseCount}">Add +1</button>
        </div>
      </div>
    `
};

define("simple-counter", SimpleCounter);
