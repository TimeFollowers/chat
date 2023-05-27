import { useDispatch, useSelector } from "react-redux";
import { decrement, increment } from "../../redux/counter/countSlice";
import { useAppSelector } from "../../redux";

function Counter() {
    const count  = useAppSelector((state) => state.counter.value)
    const dispatch = useDispatch()
    return (
        <div>
            <div>
                <button aria-label="Increment value"
                    onClick={() => dispatch(increment())}>
                    Increment
                </button>
                <span>{count}</span>
                <button aria-label="Decrement value"
                    onClick={() => dispatch(decrement())}
                >
                    Decrement
                </button>
            </div>
        </div>
    )
}

export default Counter