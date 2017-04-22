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
    "testing"
    "time"
    "encoding/json"
    "bytes"
)

func TestUnix_MarshalJSON(t *testing.T) {
    parsed, _ := time.Parse(time.RFC3339, "2017-04-22T10:11:12+00:00")
    obj := &Unix{parsed}

    marshaled, err := json.Marshal(obj)

    if err != nil {
        t.Error(err)
    }

    expected := `1492855872`
    actual := string(marshaled)

    if actual != expected {
        t.Error("Expected %s got %s", expected, actual)
    }
}

func TestUnix_UnmarshalJSON(t *testing.T) {
    str := `1492855872`

    var c Unix
    err := json.NewDecoder(bytes.NewBufferString(str)).Decode(&c)

    if err != nil {
        t.Error(err)
    }

    expected, _ := time.Parse(time.RFC3339, "2017-04-22T10:11:12+00:00")

    if c.Unix() != expected.Unix() {
        t.Errorf("Expected %d got %d", expected.Unix(), c.Unix())
    }
}

func TestUnix_UnmarshalBadJSON(t *testing.T) {
    str := `these pretzels are making me thirsty`

    var c Unix
    err := json.NewDecoder(bytes.NewBufferString(str)).Decode(&c)

    if err == nil {
        t.Error("There should be an error when unmarshaling!")
    }
}
