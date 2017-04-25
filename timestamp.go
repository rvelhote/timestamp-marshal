// Package timestamp allows JSON Marshaling and Unmarshaling of Unix Timestamps into a time.Time object
package timestamp

/*
 * The MIT License (MIT)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
import (
    "time"
    "strconv"
)

// Unix is a data type that "extends" time.Time and allows for the marshaling and unmarshaling of Unix Timestamps.
type Unix struct {
    time.Time
}

// MarshalJSON is invoked when a object is being converted to it JSON equivalent. For this data type, the marshaling
// will turn a time.Time object into its equivalent Unix Timestamp.
func (t *Unix) MarshalJSON() ([]byte, error) {
    return []byte(strconv.FormatInt(t.Time.Unix(), 10)), nil
}

// UnmarshalJSON is invoked when a JSON string representation is being converted into an object. For this data type,
// the unmarshaling will turn a Unix Timestamp into its equivalent time.Time object.
func (t *Unix) UnmarshalJSON(b []byte) error {
    timestamp, err := strconv.ParseInt(string(b), 10, 64)

    if err != nil {
        return err
    }

    t.Time = time.Unix(timestamp, 0)
    return nil
}
