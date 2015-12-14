// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"services"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void createArea(string token, Area area)")
	fmt.Fprintln(os.Stderr, "  void updateArea(string token, Area area)")
	fmt.Fprintln(os.Stderr, "  void deleteArea(string token, Area area)")
	fmt.Fprintln(os.Stderr, "   getNearBy(Coordinate coordinate, i32 limit)")
	fmt.Fprintln(os.Stderr, "   getAllAreasInCity(string cityid)")
	fmt.Fprintln(os.Stderr, "  Area getAreaById(string id)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := services.NewAreaSvcClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "createArea":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CreateArea requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg103 := flag.Arg(2)
		mbTrans104 := thrift.NewTMemoryBufferLen(len(arg103))
		defer mbTrans104.Close()
		_, err105 := mbTrans104.WriteString(arg103)
		if err105 != nil {
			Usage()
			return
		}
		factory106 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt107 := factory106.GetProtocol(mbTrans104)
		argvalue1 := services.NewArea()
		err108 := argvalue1.Read(jsProt107)
		if err108 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.CreateArea(value0, value1))
		fmt.Print("\n")
		break
	case "updateArea":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UpdateArea requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg110 := flag.Arg(2)
		mbTrans111 := thrift.NewTMemoryBufferLen(len(arg110))
		defer mbTrans111.Close()
		_, err112 := mbTrans111.WriteString(arg110)
		if err112 != nil {
			Usage()
			return
		}
		factory113 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt114 := factory113.GetProtocol(mbTrans111)
		argvalue1 := services.NewArea()
		err115 := argvalue1.Read(jsProt114)
		if err115 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.UpdateArea(value0, value1))
		fmt.Print("\n")
		break
	case "deleteArea":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeleteArea requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg117 := flag.Arg(2)
		mbTrans118 := thrift.NewTMemoryBufferLen(len(arg117))
		defer mbTrans118.Close()
		_, err119 := mbTrans118.WriteString(arg117)
		if err119 != nil {
			Usage()
			return
		}
		factory120 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt121 := factory120.GetProtocol(mbTrans118)
		argvalue1 := services.NewArea()
		err122 := argvalue1.Read(jsProt121)
		if err122 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.DeleteArea(value0, value1))
		fmt.Print("\n")
		break
	case "getNearBy":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetNearBy requires 2 args")
			flag.Usage()
		}
		arg123 := flag.Arg(1)
		mbTrans124 := thrift.NewTMemoryBufferLen(len(arg123))
		defer mbTrans124.Close()
		_, err125 := mbTrans124.WriteString(arg123)
		if err125 != nil {
			Usage()
			return
		}
		factory126 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt127 := factory126.GetProtocol(mbTrans124)
		argvalue0 := services.NewCoordinate()
		err128 := argvalue0.Read(jsProt127)
		if err128 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err129 := (strconv.Atoi(flag.Arg(2)))
		if err129 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.GetNearBy(value0, value1))
		fmt.Print("\n")
		break
	case "getAllAreasInCity":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAllAreasInCity requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetAllAreasInCity(value0))
		fmt.Print("\n")
		break
	case "getAreaById":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAreaById requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetAreaById(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
